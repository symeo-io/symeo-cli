package utils

import (
	"os"
)

func GetSystemEnv() []string {
	var env []string

	for _, e := range os.Environ() {
		env = append(env, e)
	}

	return env
}
