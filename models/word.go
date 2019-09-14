package models

import (
	"time"
	"hebrew/models/wordstate"
)

type Word struct {
	He            string          `json:"he, omitempty"`
	HeNoNikudots  string          `json:"heNoNikudots, omitempty"`
	Ru            string          `json:"ru, omitempty"`
	Transcription string          `json:"transcription, omitempty"`
	En            string          `json:"en, omitempty"`
	WordStat      *WordStatistics `json:"statistics, omitempty"`
	WordState     wordstate.State `json:"state, omitempty"`
}

type WordStatistics struct {
	WasLearned            bool
	HitCount              int
	Score                 int
	RepetitionKoefficient float32
	LastHitTime           time.Time
}
