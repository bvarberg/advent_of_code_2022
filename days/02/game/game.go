package game

type Outcome string

const (
	Win  Outcome = "Win"
	Draw Outcome = "Draw"
	Loss Outcome = "Loss"
)

type Shape string

const (
	Rock     Shape = "Rock"
	Paper    Shape = "Paper"
	Scissors Shape = "Scissors"
)
