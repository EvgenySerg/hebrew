package wordstate


type State int

const (
	New          State = 0
	Learning     State = 1
	Learnded     State = 2
	InRepetition       = 3
)
