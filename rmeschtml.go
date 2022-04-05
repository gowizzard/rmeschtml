package rmeschtml

import "strings"

// Replace is to replace the escape html characters
// The function check each escaped character & replaced them
func Replace(text string) string {

	for index, value := range list {
		text = strings.ReplaceAll(text, index, value)
	}
	return text

}
