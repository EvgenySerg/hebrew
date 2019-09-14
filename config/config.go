package config

import (
	"io/ioutil"
	"encoding/json"
)

func LoadConfig(fileName string) (*Config, error){
	c, err :=  ioutil.ReadFile(fileName)
	if err!=nil{
		return  nil,err
	}

	var config Config

	err = json.Unmarshal(c,&config)
	if err!=nil{
		return nil,err
	}
	return &config,nil
}

type Config struct {
	// When words is learned
	LearnedWordScore               int
	ScoreIncrementPerCorrectAnswer int
	ScoreDecrementByWrongAnswer    int
	ScoreDecrementByDay            int
	ScoreDecrementKoefficientByRepetition float32
}




