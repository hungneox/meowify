package cats

import (
	"math/rand"
	"strings"
)

/*
Meow main struct
*/
type Meow struct {
	Cat
}

// MeowifyLine `Meow-ify` a line
func (m Meow) MeowifyLine(line string) string {
	return m.Map(line, m.Meowify)
}

// Meowify `Meow-ify` a word
func (m Meow) Meowify(word string) string {
	length := len(word)

	switch length {
	case 1:
		return m.ToProperCase("m", word)
	case 2:
		return m.ToProperCase("me", word)
	case 3:
		return m.ToProperCase("meo", word)
	case 4:
		return m.ToProperCase("meow", word)
	}

	remain := int(length - len("meow"))

	//TODO Handle word is longer than 5 chars
	if remain >= 1 {

		max := remain - 1

		if max < 1 {
			max = 1
		}
		// random amount of eeee
		es := rand.Intn(max) + 1

		// some oooo left
		os := remain - es

		return m.ToProperCase("me", word) + strings.Repeat("e", es) + strings.Repeat("o", os) + "ow"
	}

	return word
}
