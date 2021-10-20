package helpers

import (
	"path"
	"runtime"
)

func EnvPath() [2]string {
	_, configPath, _, _ := runtime.Caller(0)
	pathPrefix := path.Join(path.Dir(configPath), "..")
	mockedEnvPath := path.Join(pathPrefix, "./config/mocks/.env")
	envPath := path.Join(pathPrefix, "./.env")
	return [2]string{mockedEnvPath, envPath}
}
