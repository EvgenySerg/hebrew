package models


func NewWordSet(name string) *WordSet{
	words:=make([]*Word,0)
	return &WordSet{Name:name, Words:words}
}

type WordSet struct {
	Name string
	Words []*Word
}
