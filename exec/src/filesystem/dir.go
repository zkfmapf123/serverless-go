package filesystem

import (
	"os"
	"strings"

	"github.com/zkfmapf123/serverless-go-deploy-agent/src/interaction"
	"github.com/zkfmapf123/serverless-go-deploy-agent/src/utils"
)

func SelectBoxDirectory(path string) (string, bool) {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	newPath := utils.Concat(p, "/", path)

	entities, err := os.ReadDir(newPath)
	if err != nil {
		panic(err)
	}

	var dirs []string
	for _, v := range entities {
		dirs = append(dirs, v.Name())
	}

	entity, isExit := interaction.Select("Select Function", dirs)

	if isExit {
		return "", true
	}

	newPath = utils.Concat(newPath, "/", entity)

	return newPath, false
}

func IsExist(path, name string) bool {

	entities, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, v := range entities {
		if strings.Contains(v.Name(), name) {
			return true
		}
	}

	return false
}

func MakeZip(path string) {
	err := os.Chdir(path)
	if err != nil {
		panic(err)
	}

	envMap := map[string]string{
		"GOOS":        "linux",
		"GOARCH":      "arm64",
		"CGO_ENABLED": "0",
	}

	for k, v := range envMap {
		os.Setenv(k, v)
	}

	interaction.Exec("rm", "-rf", "bootstrap", "*.zip")
	interaction.Exec("go", "build", "-tags", "lambda.norpc", "-ldflags", "-s -w", "-o", "bootstrap", "main.go")
	interaction.Exec("chmod", "+x", "bootstrap")
	interaction.Exec("zip", "bootstrap.zip", "bootstrap")
}
