package interaction

import (
	"github.com/AlecAivazis/survey/v2"
)

func Select(msg string, list []string) (string, bool) {
	var answer string

	prompt := &survey.Select{
		Message: msg,
		Options: append(list, "exit..."),
	}
	survey.AskOne(prompt, &answer)

	if answer == "exit..." {
		return "", true
	}

	return answer, false
}
