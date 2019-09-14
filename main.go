package main

import (
	ui2 "hebrew/ui"
	"hebrew/models"
	"hebrew/config"
	"log"
)

func main() {

	var dict models.Dict

	err := dict.Load("dict.json")
	if err != nil {
		log.Panic(err)
	}

	cfg,err:=config.LoadConfig("config.json")
	if err!=nil{
		log.Panic(err.Error())
	}


	ui:=ui2.NewUi(cfg, &dict)
	ui.Show()


	//
	//
	//wordSet := dict.GetAllWordsWordSet()
	//man:= manager2.Manager{DesiredCorrectCount:10}
	//quiz:=manager2.NewQuiz(wordSet)
	//man.StartQuiz(quiz)
	//
	//Fin()
}

