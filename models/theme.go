package models

import "hebrew/models/wordstate"

type Theme struct {
	Name    string
	Words   []*Word
	learned int
}

func (th *Theme) GetWordsByState(state wordstate.State) []*Word {
	var words []*Word
	for _, word := range th.Words {
		if word.WordState == state {
			words = append(words, word)
		}
	}
	return words
}
