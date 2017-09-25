package cats

import (
	"regexp"
	"strings"
	"unicode"
)

/*
Cat main struct
*/
type Cat struct {
	CatInterface
}

//FindWords return proper word in given token (excludes punctuation)
func (c Cat) FindWords(token string) []string {
	r, _ := regexp.Compile("\\w+")
	return r.FindAllString(token, -1)
}

//ToProperCase Convert potato word to proper case
func (c Cat) ToProperCase(word string, original string) string {
	originalRunes := []rune(original)
	MeowRunes := []rune(word)

	for i := 0; i < len(MeowRunes); i++ {
		if unicode.IsUpper(originalRunes[i]) {
			MeowRunes[i] = unicode.ToUpper(MeowRunes[i])
		} else if unicode.IsLower(originalRunes[i]) {
			MeowRunes[i] = unicode.ToLower(MeowRunes[i])
		}
	}

	return string(MeowRunes)
}

// Map function to each line
func (c Cat) Map(line string, meowify func(string) string) string {
	words := strings.Fields(line)

	// Replace each word with potato one
	for i := 0; i < len(words); i++ {
		var tokens = c.FindWords(words[i])
		for j := 0; j < len(tokens); j++ {
			words[i] = strings.Replace(words[i], tokens[j], meowify(tokens[j]), -1)
		}
	}

	return strings.Join(words, " ")
}
