package camGRpcClient

import (
	"google.golang.org/grpc"
	"sync"
)

type connDict struct {
	sync.Map
}

func (dict *connDict) Set(key string, conn *grpc.ClientConn) {
	dict.Store(key, conn)
}

func (dict *connDict) Get(key string) (*grpc.ClientConn, bool) {
	connI, has := dict.Load(key)
	if has {
		var conn *grpc.ClientConn
		if connI != nil {
			conn = connI.(*grpc.ClientConn)
		}
		return conn, true
	}
	return nil, false
}

func (dict *connDict) Del(key string) {
	dict.Delete(key)
}

func (dict *connDict) Range(handler func(key string, conn *grpc.ClientConn) bool) {
	dict.Map.Range(func(key, value interface{}) bool {
		keyStr := key.(string)
		valuePtr := value.(*grpc.ClientConn)
		return handler(keyStr, valuePtr)
	})
}

func (dict *connDict) Len() int {
	l := 0
	dict.Range(func(key string, conn *grpc.ClientConn) bool {
		l++
		return true
	})
	return l
}
