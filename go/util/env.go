package util

import (
	"os"
	"strconv"
)

const EnvKeyIsOnDocker = "IS_ON_DOCKER"

// Docker 上で実行されているかどうか
func IsOnDocker() bool {
	value := os.Getenv(EnvKeyIsOnDocker)
	if len(value) == 0{
		return false
	}

	isOnDocker, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return isOnDocker
}
