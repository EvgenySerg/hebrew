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
	dict.Name = "Mathematics"
	dict.Themes = []*models.Theme{
		{Name: "Сложение", Words: []*models.Word{
			{Ru: "сложение", He: "חִיבּוּר", HeNoNikudots: "חיבור",},
			{Ru: "сложить", He: "לְחַבֵּר", HeNoNikudots: "לחבר",},
			{Ru: "слагаемое", He: "מְחוּבָּר", HeNoNikudots: "מחובר",},
			{Ru: "сумма", He: "סכוּם, סָך", HeNoNikudots: "סכום, סך",},
		},
		},
		{Name: "Вычитание", Words: []*models.Word{
			{Ru: "вычитание", He: "חִיסוּר", HeNoNikudots: "חיסור",},
			{Ru: "вычитаемое", He: "מְחַסֵר", HeNoNikudots: "מחסר",},
			{Ru: "разность", He: "הֶפרֵש", HeNoNikudots: "הפרש",},
			{Ru: "отними", He: "הַפְחֵת", HeNoNikudots: "הפחת",},
		},
		},

	}



	err := dict.Load("dict.json")
	if err != nil {
		fmt.Println(err.Error())
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