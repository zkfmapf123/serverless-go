package filesystem

import (
	"os"

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
