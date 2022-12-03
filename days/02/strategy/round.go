package strategy

type Round struct {
	opponentThrow Shape
	throw         Shape
}

func NewRound(opponentThrow Shape, throw Shape) Round {
	return Round{
		opponentThrow,
		throw,
	}
}

func (r Round) Play() Outcome {
	switch r.opponentThrow {
	case Rock:
		switch r.throw {
		case Rock:
			return Draw
		case Paper:
			return Win
		case Scissors:
			return Loss
		}
	case Paper:
		switch r.throw {
		case Rock:
			return Loss
		case Paper:
			return Draw
		case Scissors:
			return Win
		}
	case Scissors:
		switch r.throw {
		case Rock:
			return Win
		case Paper:
			return Loss
		case Scissors:
			return Draw
		}
	}

	panic("Invalid throws")
}

func (r Round) Score() int {
	outcome := r.Play()
	return ScoreOutcome(outcome) + ScoreThrow(r.throw)
}
