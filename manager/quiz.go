package manager

import "hebrew/models"

func NewQuiz(set *models.WordSet) *Quiz{
	scoreTable := make ([]*models.ScoreUnit,0)
	for _,word:=range set.Words{
		scoreUnit:= models.ScoreUnit{
			Word:       word,
			K:          0,
			OkCount:    0,
			TotalCount: 0,
		}
		scoreTable = append(scoreTable, &scoreUnit)
	}
	return &Quiz {WordSet: set, currentIndex:0,ScoreTable:scoreTable}
}

type Quiz struct {
	currentIndex int
	ScoreTable   []*models.ScoreUnit
	WordSet      *models.WordSet
}

func (q *Quiz) GetLearnedWordCount() int{
	count := 0
	for _, scu:=range q.ScoreTable {
		if scu.IsLearned{
			count++
		}
	}
	return count
}

func (q *Quiz) GetWordScore(word *models.Word) *models.ScoreUnit{
	for _, scu := range q.ScoreTable{
		if scu.Word.Ru==word.Ru{
			return scu
		}
	}
	return nil
}

func (q *Quiz) NextWord() *models.Word{
	q.currentIndex++
	if q.currentIndex>=len(q.WordSet.Words)-1{
		q.currentIndex=0
	}
	return q.WordSet.Words[q.currentIndex]
}
