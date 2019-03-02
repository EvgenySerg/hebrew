package manager

import (
	"hebrew/models"
	"bufio"
	"fmt"
	"os"
)

type Manager struct {
	DesiredCorrectCount int
	CurrentCorrectCount int
	Quiz                *Quiz

}

func (man *Manager)StartQuiz(quiz *Quiz){

	man.DesiredCorrectCount =10
	man.Quiz=quiz
	learned:=man.Quiz.GetLearnedWordCount()
	scoreTableCount := len(man.Quiz.ScoreTable)
	for learned!= scoreTableCount{
		nextWord:=quiz.NextWord()
		scu, isCorrect := man.ProcessHeToRuWord(nextWord)
		percentage := man.CurrentCorrectCount/man.DesiredCorrectCount
		if percentage >= 1{
			scu.IsLearned=true
		}

		if isCorrect{
			fmt.Printf("Correct %d", percentage)
		} else {
			fmt.Printf("Incorrect %d", percentage)
		}
		learned=man.Quiz.GetLearnedWordCount()
	}
}

func (man *Manager) ProcessHeToRuWord(word *models.Word) (*models.ScoreUnit, bool){
	answer:= man.ShowWord(word.He)
	scu:=man.Quiz.GetWordScore(word)
	correct:=false
	if answer==word.Ru {
		correct = true
		scu.OkCount++
	}
	scu.TotalCount++
	scu.K = (float32(scu.OkCount)/ float32(man.DesiredCorrectCount))/float32(scu.TotalCount)
	return scu, correct
}

func (man *Manager) ShowWord (word string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nWord is: \n" + word)
	answer, _ := reader.ReadString('\n')
	return answer[:len(answer)-1]
}