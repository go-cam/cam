package camCache

import (
	"errors"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camUtils"
	"time"
)

type FileCache struct {
	CacheInterface

	DirDepth   int           // cache dir depth
	FilePath   string        // cache dir
	GcCheckNum int           // number of times to check file whether GC
	GCInterval time.Duration // GC interval
}

func NewFileCache() *FileCache {
	cache := new(FileCache)
	cache.DirDepth = 2
	cache.FilePath = camUtils.File.GetRunPath() + "/runtime/cache"
	cache.GcCheckNum = 1000
	cache.GCInterval = 6 * time.Hour
	return cache
}

// init engine
func (cache *FileCache) Init() error {
	if cache.DirDepth > 8 {
		return errors.New("DirDepth cannot be greater than 8. ")
	}
	if cache.FilePath == "" {
		return errors.New("FilePath cannot be empty. ")
	}
	return nil
}

// put key-value to engine
func (cache *FileCache) Set(key string, value interface{}, duration time.Duration) error {
	filename := cache.getCacheFile(key)
	ao := NewFileCacheAo(value, duration)
	bytes := camUtils.Json.Encode(ao)

	var err error
	dir := camUtils.File.Dir(filename)
	if !camUtils.File.Exists(dir) {
		err = camUtils.File.Mkdir(dir)
		if err != nil {
			return err
		}
	}

	return camUtils.File.WriteFile(filename, bytes)
}

// get value
func (cache *FileCache) Get(key string) interface{} {
	filename := cache.getCacheFile(key)
	if !camUtils.File.Exists(filename) {
		return nil
	}

	ao := cache.getFileCacheAo(filename)
	if ao == nil {
		return nil
	}

	if ao.DeadTimestamp < float64(time.Now().Unix()) {
		depth := 1
		err := cache.gcSubPath(filename, &depth)
		if err != nil {
			camStatics.App.Error("FileCache.Get", err.Error())
		}
		return nil
	}

	return ao.Value
}

// delete value form engine
func (cache *FileCache) Del(keys ...string) error {
	var err error
	for _, key := range keys {
		err = cache.delOne(key)
		if err != nil {
			return err
		}
	}
	return nil
}

// delete one value
func (cache *FileCache) delOne(key string) error {
	filename := cache.getCacheFile(key)
	if !camUtils.File.Exists(filename) {
		return nil
	}
	return camUtils.File.Remove(filename)
}

// get cache save absolute filename
func (cache *FileCache) getCacheFile(key string) string {
	filename := camUtils.Encrypt.Md5([]byte(key))
	dirname := ""
	for i := 0; i < cache.DirDepth; i++ {
		if i > 8 {
			break
		}
		start := i * 2
		end := start + 2
		dirname = dirname + "/" + filename[start:end]
	}

	return cache.FilePath + dirname + "/" + filename + ".bin"
}

// garbage collection
func (cache *FileCache) GC() error {
	gcNum := cache.GcCheckNum
	return cache.gcSubPath(cache.FilePath, &gcNum)
}

// garbage collection sub file or dir
func (cache *FileCache) gcSubPath(parentPath string, gcNum *int) error {
	if *gcNum < 0 {
		return nil
	}
	if !camUtils.File.Exists(parentPath) {
		return nil
	}

	if !camUtils.File.IsDir(parentPath) {
		*gcNum--
		ao := cache.getFileCacheAo(parentPath)
		if ao.IsDeadNano() {
			return camUtils.File.Remove(parentPath)
		}
		return nil
	} else {
		fileInfoList, err := camUtils.File.ScanDir(parentPath, true)
		if err != nil {
			return err
		}

		for _, fileInfo := range fileInfoList {
			path := parentPath + "/" + fileInfo.Name()
			err = cache.gcSubPath(path, gcNum)
			if err != nil {
				return err
			}
		}

		// check folder whether empty
		fileInfoList, err = camUtils.File.ScanDir(parentPath, true)
		if err != nil {
			return err
		}
		if len(fileInfoList) == 0 && parentPath != cache.FilePath {
			return camUtils.File.Remove(parentPath)
		}

		return nil
	}
}

// clear all cache
func (cache *FileCache) Flush() error {
	fileInfoList, err := camUtils.File.ScanDir(cache.FilePath, true)
	if err != nil {
		return err
	}

	for _, fileInfo := range fileInfoList {
		if !fileInfo.IsDir() {
			continue
		}

		path := cache.FilePath + "/" + fileInfo.Name()
		err = camUtils.File.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	return nil
}

// get FileCacheAo
func (cache *FileCache) getFileCacheAo(filename string) *FileCacheAo {
	bytes, err := camUtils.File.ReadFile(filename)
	if err != nil {
		camStatics.App.Error("FileCache.Get", err.Error())
		return nil
	}

	ao := new(FileCacheAo)
	camUtils.Json.DecodeToObj(bytes, ao)
	return ao
}

// get GC interval
func (cache *FileCache) GetGCInterval() time.Duration {
	return cache.GCInterval
}
