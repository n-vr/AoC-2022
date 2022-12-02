// Package rps implements rock, paper, scissors,
// with the rules according to AoC day 2 challenge.
package rps

// A shape can be a Rock, Paper or Scissors.
type Shape int

// Constants for which shape can be chosen.
const (
	ShapeRock Shape = iota
	ShapePaper
	ShapeScissors
)

// A player can be Me or the Opponent.
type Player int

// Constants for different players.
const (
	PlayerMe = iota
	PlayerOpponent
)

// An outcome can be Win, Tie or Lose.
type Outcome rune

// Constants for the different outcomes.
const (
	OutcomeLose Outcome = 'X'
	OutcomeTie  Outcome = 'Y'
	OutcomeWin  Outcome = 'Z'
)

// Constants for the opponents rock, paper or scissors.
const (
	OpponentRock     = 'A'
	OpponentPaper    = 'B'
	OpponentScissors = 'C'
)

// Constants for my opponents rock, paper or scissors.
const (
	MyRock     = 'X'
	MyPaper    = 'Y'
	MyScissors = 'Z'
)

// Constants for the scores you get for a chosen shape.
const (
	ScoreRock     = 1
	ScorePaper    = 2
	ScoreScissors = 3
)

// Constants for the scores you get for the round outcome.
const (
	ScoreLost = 0
	ScoreTied = 3
	ScoreWon  = 6
)

// NewShape creates a new shape from a rune.
func NewShape(r rune) *Shape {
	meShapes := map[rune]Shape{
		MyRock:     ShapeRock,
		MyPaper:    ShapePaper,
		MyScissors: ShapeScissors,
	}

	opponentShapes := map[rune]Shape{
		OpponentRock:     ShapeRock,
		OpponentPaper:    ShapePaper,
		OpponentScissors: ShapeScissors,
	}

	shape, ok := meShapes[r]
	if !ok {
		shape = opponentShapes[r]
	}

	return &shape
}

// Rune returns the rune that corresponds to the shape and the player.
func (s *Shape) Rune(player Player) rune {
	meShapes := map[Shape]rune{
		ShapeRock:     MyRock,
		ShapePaper:    MyPaper,
		ShapeScissors: MyScissors,
	}

	opponentShapes := map[Shape]rune{
		ShapeRock:     OpponentRock,
		ShapePaper:    OpponentPaper,
		ShapeScissors: OpponentScissors,
	}

	if player == PlayerMe {
		return meShapes[*s]
	}
	return opponentShapes[*s]
}

// RoundScore returns the total score
// from the shape used and the round outcome.
func RoundScore(opponent, me Shape) int {
	return ShapeScore(me) + OutcomeScore(opponent, me)
}

// ShapeScore returns the score you got from the shape used.
func ShapeScore(shape Shape) int {
	shapeScores := map[rune]int{
		MyRock:     ScoreRock,
		MyPaper:    ScorePaper,
		MyScissors: ScoreScissors,
	}

	return shapeScores[shape.Rune(PlayerMe)]
}

// OutcomeScore returns the score you got from the round outcome.
func OutcomeScore(opponent, me Shape) int {
	// Tied
	if opponent == me {
		return ScoreTied
	}

	// Won
	if me == opponent+1 || me == ShapeRock && opponent == ShapeScissors {
		return ScoreWon
	}

	// Lost
	return ScoreLost
}

// ChooseShapeForOutcome return the shape you need to choose
// to achieve the desired outcome, based on the opponents [Shape].
func ChooseShapeForOutcome(opponent Shape, outcome Outcome) Shape {
	switch outcome {
	case OutcomeTie:
		return opponent
	case OutcomeLose:
		if opponent == ShapeRock {
			return ShapeScissors
		}
		return opponent - 1
	default:
		// Default is to win ofcourse :).
		if opponent == ShapeScissors {
			return ShapeRock
		}
		return opponent + 1
	}
}
