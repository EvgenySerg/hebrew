package models

type ScoreUnit struct {
	Word       *Word
	OkCount    int
	TotalCount int
	K          float32
	IsLearned  bool
}
