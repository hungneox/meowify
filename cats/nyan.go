package cats

/*
Nyan (ニャンニャン) main struct
*/
type Nyan struct {
	Cat
}

//MeowifyLine each line
func (n Nyan) MeowifyLine(line string) string {
	return n.Map(line, n.Meowify)
}

// Meowify a word
func (n Nyan) Meowify(word string) string {
	length := len(word)

	switch length {
	case 1:
		return n.ToProperCase("n", word)
	case 2:
		return n.ToProperCase("ny", word)
	case 3:
		return n.ToProperCase("nya", word)
	case 4:
		return n.ToProperCase("nyan", word)
	}

	left := length - len("nyan")

	//TODO Handle word is longer than 5 chars
	if left > 0 {
		return n.ToProperCase("nyan", word) + n.Meowify(word[length-left:])
	}

	return word
}
