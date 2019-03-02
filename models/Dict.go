package models

import (
	"io/ioutil"
	"encoding/json"
	"os"
)

type Dict struct {
	Name string
	Themes []*Theme
}


func (dict *Dict) GetAllWordsWordSet() *WordSet{
	wordSet:= NewWordSet("All words")

	for _, theme :=range dict.Themes{
		wordSet.Words = append(wordSet.Words, theme.Words...)
	}
	return wordSet
}


func (dict *Dict) Load(filename string)  error{
	words, err :=  ioutil.ReadFile(filename)
	if err!=nil{
		return  err
	}

	err = json.Unmarshal(words, &dict)
	if err!=nil{
		return err
	}
	return  nil
}


func (dict *Dict) Save(filename string) error{
	data, err:=json.Marshal(dict)
	if err!=nil{
		return err
	}
	err = ioutil.WriteFile(filename, data, os.ModeAppend)
	if err!=nil{
		return err
	}
	return nil
}



