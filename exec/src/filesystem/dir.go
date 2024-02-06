package filesystem

import (
	"fmt"
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

func InjectFileScript(path, filename, txt string) error {
	file, err := os.Create(fmt.Sprintf("%s/%s", path, filename))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(txt)
	if err != nil {
		return err
	}

	return nil
}
