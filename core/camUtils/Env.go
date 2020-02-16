package camUtils

import "strings"

type EnvUtil struct {
	dict map[string]string
}

var Env = newEnvUtil()

func newEnvUtil() *EnvUtil {
	env := new(EnvUtil)
	env.autoLoadEnvDict()

	return env
}

func (util *EnvUtil) Get(key string) string {
	value, has := util.dict[key]
	if has {
		return value
	}
	return ""
}

func (util *EnvUtil) autoLoadEnvDict() {
	util.dict = map[string]string{}

	envFilename := File.GetRunPath() + "/.env"
	if !File.Exists(envFilename) {
		return
	}

	contentBytes, err := File.ReadFile(envFilename)
	Error.Panic(err)

	contentStr := string(contentBytes)
	lines := strings.Split(contentStr, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "#") ||
			strings.HasPrefix(line, "//") {
			continue
		}

		equalKey := strings.Index(line, "=")
		if equalKey == -1 {
			continue
		}

		key := strings.TrimSpace(line[0:equalKey])
		value := strings.TrimSpace(line[equalKey+1:])

		util.dict[key] = value
	}
}
