package main

import (
	"testing"
	"learn/tic/players"
)

func TestReadArgs(t *testing.T) {

	cases := []struct {
		description       string
		expectedBoardSize int
		player1Sign       string
		player2Sign       string
		computerSign      string
	}{
		{
			description:       "Should return given size as board size without error",
			expectedBoardSize: 10,
			player2Sign:"r",
			player1Sign:"j",
			computerSign:"*",
		},
	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {

			readArgs("test_config.txt")

			if c.expectedBoardSize != boardSize {
				t.Errorf("\nUnexpected board size: \n\t\t expected: %d, \n\t\t actual: %d", c.expectedBoardSize, boardSize)
			}

			for i:= range players.GetPlayers() {
				if !(players.GetPlayers()[i].GetSign()==c.player1Sign || players.GetPlayers()[i].GetSign()==c.player2Sign || players.GetPlayers()[i].GetSign()==c.computerSign) {
					t.Errorf("\n Unexpected signs of players \n")
				}
			}
		})
	}
}

