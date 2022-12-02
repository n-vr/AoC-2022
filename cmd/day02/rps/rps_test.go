package rps

import "testing"

func TestNewShape(t *testing.T) {
	cases := []struct {
		desc  string
		input rune
		want  Shape
	}{
		{"OpponentRock", 'A', ShapeRock},
		{"OpponentPaper", 'B', ShapePaper},
		{"OpponentScissors", 'C', ShapeScissors},
		{"MyRock", 'X', ShapeRock},
		{"MyPaper", 'Y', ShapePaper},
		{"MyScissors", 'Z', ShapeScissors},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(tt *testing.T) {
			got := NewShape(tc.input)
			if tc.want != *got {
				t.Errorf("Want: %v, Got: %v", tc.want, got)
			}
		})
	}
}

func TestRune(t *testing.T) {
	cases := []struct {
		desc        string
		inputShape  Shape
		inputPlayer Player
		want        rune
	}{
		{"OpponentRock", ShapeRock, PlayerOpponent, 'A'},
		{"OpponentPaper", ShapePaper, PlayerOpponent, 'B'},
		{"OpponentScissors", ShapeScissors, PlayerOpponent, 'C'},
		{"MyRock", ShapeRock, PlayerMe, 'X'},
		{"MyPaper", ShapePaper, PlayerMe, 'Y'},
		{"MyScissors", ShapeScissors, PlayerMe, 'Z'},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(tt *testing.T) {
			got := tc.inputShape.Rune(tc.inputPlayer)
			if tc.want != got {
				t.Errorf("Want: %c, Got: %c", tc.want, got)
			}
		})
	}
}

func TestRoundScore(t *testing.T) {
	cases := []struct {
		desc          string
		inputOpponent Shape
		inputMe       Shape
		want          int
	}{
		{"RockRockTie", ShapeRock, ShapeRock, ScoreTied + ScoreRock},
		{"RockPaperWin", ShapeRock, ShapePaper, ScoreWon + ScorePaper},
		{"RockScissorsLose", ShapeRock, ShapeScissors, ScoreLost + ScoreScissors},
		{"PaperRockLose", ShapePaper, ShapeRock, ScoreLost + ScoreRock},
		{"PaperPaperTie", ShapePaper, ShapePaper, ScoreTied + ScorePaper},
		{"PaperScissorsWin", ShapePaper, ShapeScissors, ScoreWon + ScoreScissors},
		{"ScissorsRockWin", ShapeScissors, ShapeRock, ScoreWon + ScoreRock},
		{"ScissorsPaperLose", ShapeScissors, ShapePaper, ScoreLost + ScorePaper},
		{"ScissorsScissorsTie", ShapeScissors, ShapeScissors, ScoreTied + ScoreScissors},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(tt *testing.T) {
			got := RoundScore(tc.inputOpponent, tc.inputMe)
			if tc.want != got {
				t.Errorf("Want: %v, Got: %v", tc.want, got)
			}
		})
	}
}

func TestShapeScore(t *testing.T) {
	cases := []struct {
		desc  string
		input Shape
		want  int
	}{
		{"Rock", ShapeRock, ScoreRock},
		{"Paper", ShapePaper, ScorePaper},
		{"Scissors", ShapeScissors, ScoreScissors},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(tt *testing.T) {
			got := ShapeScore(tc.input)
			if tc.want != got {
				t.Errorf("Want: %v, Got: %v", tc.want, got)
			}
		})
	}
}

func TestOutcomeScore(t *testing.T) {
	cases := []struct {
		desc          string
		inputOpponent Shape
		inputMe       Shape
		want          int
	}{
		{"RockRockTie", ShapeRock, ShapeRock, ScoreTied},
		{"RockPaperWin", ShapeRock, ShapePaper, ScoreWon},
		{"RockScissorsLose", ShapeRock, ShapeScissors, ScoreLost},
		{"PaperRockLose", ShapePaper, ShapeRock, ScoreLost},
		{"PaperPaperTie", ShapePaper, ShapePaper, ScoreTied},
		{"PaperScissorsWin", ShapePaper, ShapeScissors, ScoreWon},
		{"ScissorsRockWin", ShapeScissors, ShapeRock, ScoreWon},
		{"ScissorsPaperLose", ShapeScissors, ShapePaper, ScoreLost},
		{"ScissorsScissorsTie", ShapeScissors, ShapeScissors, ScoreTied},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(tt *testing.T) {
			got := OutcomeScore(tc.inputOpponent, tc.inputMe)
			if tc.want != got {
				t.Errorf("Want: %v, Got: %v", tc.want, got)
			}
		})
	}
}

func TestChooseShapeForOutcome(t *testing.T) {
	cases := []struct {
		desc          string
		inputOpponent Shape
		inputOutcome  Outcome
		want          Shape
	}{
		{"RockLoseScissors", ShapeRock, OutcomeLose, ShapeScissors},
		{"RockTieRock", ShapeRock, OutcomeTie, ShapeRock},
		{"RockWinPaper", ShapeRock, OutcomeWin, ShapePaper},
		{"PaperLoseRock", ShapePaper, OutcomeLose, ShapeRock},
		{"PaperTiePaper", ShapePaper, OutcomeTie, ShapePaper},
		{"PaperWinScissors", ShapePaper, OutcomeWin, ShapeScissors},
		{"ScissorsLosePaper", ShapeScissors, OutcomeLose, ShapePaper},
		{"ScissorsTieScissors", ShapeScissors, OutcomeTie, ShapeScissors},
		{"ScissorsWinRock", ShapeScissors, OutcomeWin, ShapeRock},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(tt *testing.T) {
			got := ChooseShapeForOutcome(tc.inputOpponent, tc.inputOutcome)
			if tc.want != got {
				t.Errorf("Want: %c, Got: %c", tc.want.Rune(PlayerMe), got.Rune(PlayerMe))
			}
		})
	}
}
