package main

import (
	"hebrew/models"
	"fmt"
	"os"
	"bufio"
	manager2 "hebrew/manager"
)

func main() {
	var dict models.Dict

	err := dict.Load("dict.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}


	wordSet := dict.GetAllWordsWordSet()
	man:= manager2.Manager{DesiredCorrectCount:10}
	quiz:=manager2.NewQuiz(wordSet)
	man.StartQuiz(quiz)

	Fin()
}


func Fin(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press any key to Exit")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}