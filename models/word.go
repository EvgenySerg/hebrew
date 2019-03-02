package models

type Word struct {
	He            string `json:"he, omitempty"`
	HeNoNikudots  string `json:"heNoNikudots, omitempty"`
	Ru            string `json:"ru, omitempty"`
	Transcription string `json:"transcription, omitempty"`
	En            string `json:"en, omitempty"`
}
