package common

import "regexp"

type Params map[string]any

func IsWord(content string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]+(-[a-zA-Z]+)*$`)
	return re.Match([]byte(content))
}
