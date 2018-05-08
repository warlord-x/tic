package main

import (
	"testing"
)


func TestCheckRowColumnDiagWin(t *testing.T) {

	cases := []struct {
		description       string
		board             [][]string
		expectedResultRow int
		expectedResultColumn int
		expectedResultDiag int
	}{
		{
			description:       "should return zero",
			board:             [][]string{{"_", "_", "_", "_"}, {"_", "_", "_", "_"}, {"_", "_", "_", "_"}, {"_", "_", "_", "_"}},
			expectedResultRow: 0,
			expectedResultColumn:0,
			expectedResultDiag:0,
		},
		{
			description:       "should return opponent wins",
			board:             [][]string{{"x", "x", "o"}, {"o", "o", "o"}, {"x", "x", "_"}},
			expectedResultRow: -1,
			expectedResultColumn:0,
			expectedResultDiag:0,
		},
		{
			description:       "should return my row 3 win",
			board:             [][]string{{"_", "o", "o"}, {"x", "o", "o"}, {"x", "x", "x"}},
			expectedResultRow: 1,
			expectedResultColumn:0,
			expectedResultDiag:0,
		},
		{
			description:       "should return my row 1 win",
			board:             [][]string{{"x", "x", "x"}, {"o", "o", "x"}, {"o", "o", "_"}},
			expectedResultRow: 1,
			expectedResultColumn:0,
			expectedResultDiag:0,
		},
		{
			description:       "should return my row 2 win",
			board:             [][]string{{"o", "o", "x"}, {"x", "x", "x"}, {"o", "o", "_"}},
			expectedResultRow: 1,
			expectedResultColumn:0,
			expectedResultDiag:0,
		},
		{
			description:       "should return opponent col 1 win",
			board:             [][]string{{"o", "x", "x"}, {"o", "x", "o"}, {"o", "o", "_"}},
			expectedResultRow: 0,
			expectedResultColumn:-1,
			expectedResultDiag:0,
		},
		{
			description:       "should return opponent col 2 win",
			board:             [][]string{{"x", "o", "x"}, {"o", "o", "x"}, {"o", "o", "_"}},
			expectedResultRow: 0,
			expectedResultColumn:-1,
			expectedResultDiag:0,
		},
		{
			description:       "should return my col 3 win",
			board:             [][]string{{"o", "o", "x"}, {"x", "o", "x"}, {"o", "x", "x"}},
			expectedResultRow: 0,
			expectedResultColumn:1,
			expectedResultDiag:0,
		},
		{
			description:       "should return my col 1 win",
			board:             [][]string{{"x", "o", "x"}, {"x", "o", "x"}, {"x", "x", "o"}},
			expectedResultRow: 0,
			expectedResultColumn:1,
			expectedResultDiag:0,
		},
		{
			description:       "should return my col 2 win",
			board:             [][]string{{"o", "x", "x"}, {"x", "x", "o"}, {"o", "x", "x"}},
			expectedResultRow: 0,
			expectedResultColumn:1,
			expectedResultDiag:0,
		},
		{
			description:       "should return opponent col 3 win",
			board:             [][]string{{"o", "x", "o"}, {"x", "x", "o"}, {"x", "o", "o"}},
			expectedResultRow: 0,
			expectedResultColumn:-1,
			expectedResultDiag:0,
		},
		{
			description:       "should return opponent diag right win",
			board:             [][]string{{"o", "x", "o"}, {"x", "o", "o"}, {"o", "x", "x"}},
			expectedResultRow: 0,
			expectedResultColumn:0,
			expectedResultDiag:-1,
		},
		{
			description:       "should return opponent diag left win",
			board:             [][]string{{"o", "x", "o"}, {"x", "o", "x"}, {"x", "o", "o"}},
			expectedResultRow: 0,
			expectedResultColumn:0,
			expectedResultDiag:-1,
		},
		{
			description:       "should return my diag right win",
			board:             [][]string{{"o", "x", "x"}, {"x", "x", "o"}, {"x", "o", "o"}},
			expectedResultRow: 0,
			expectedResultColumn:0,
			expectedResultDiag:1,
		},
		{
			description:       "should return my diag left win",
			board:             [][]string{{"x", "x", "o"}, {"x", "x", "o"}, {"o", "o", "x"}},
			expectedResultRow: 0,
			expectedResultColumn:0,
			expectedResultDiag:1,
		},
		{
			description:       "should return opponent win with 2 opponents for row",
			board:             [][]string{{"x", "x", "o"}, {"z", "z", "z"}, {"o", "o", "x"}},
			expectedResultRow: -1,
			expectedResultColumn:0,
			expectedResultDiag:0,
		},
		{
			description:       "should return opponent win with 2 opponents for col",
			board:             [][]string{	{"z", "x", "o"},
											{"z", "o", "x"},
											{"z", "o", "x"}},
			expectedResultRow: 0,
			expectedResultColumn:-1,
			expectedResultDiag:0,
		},



	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			computerSign ="x"
			opponent1Sign="o"
			opponent2Sign="z"
			resultRow := checkRow(c.board)
			resultColumn:=checkColumn(c.board)
			resultDiag:=checkDiag(c.board)

			if c.expectedResultRow != resultRow {
				t.Errorf("\nUnexpected resultRow: \n\t\t expected: %d, \n\t\t actual: %d", c.expectedResultRow, resultRow)

			}
			if c.expectedResultColumn != resultColumn {
				t.Errorf("\nUnexpected resultColumn : \n\t\t expected: %d, \n\t\t actual: %d", c.expectedResultColumn, resultColumn)

			}
			if c.expectedResultDiag != resultDiag {
				t.Errorf("\nUnexpected resultDiag : \n\t\t expected: %d, \n\t\t actual: %d", c.expectedResultDiag, resultDiag)

			}
		})
	}
}

func TestWeightMove(t *testing.T) {

	cases := []struct {
		description       string
		board             [][]string
		expectedMoveWeight int
	}{
		{
			description:       "no one wins",
			board:             [][]string{{"_", "_", "_", "_"}, {"_", "_", "_", "_"}, {"_", "_", "_", "_"}, {"_", "_", "_", "_"}},
			expectedMoveWeight: 0,
		},
		{
			description:       "my win by row",
			board:             [][]string{{"x", "x", "x"}, {"o", "o", "x"}, {"o", "o", "_"}},
			expectedMoveWeight: 1,
		},
		{
			description:       "opposite win by row",
			board:             [][]string{{"x", "x", "o"}, {"o", "o", "o"}, {"x", "x", "_"}},
			expectedMoveWeight: -1,
		},
		{
			description:       "my win by column",
			board:             [][]string{{"o", "x", "x"}, {"x", "x", "o"}, {"o", "x", "x"}},
			expectedMoveWeight: 1,
		},
		{
			description:       "opposite win by column",
			board:             [][]string{{"o", "x", "o"}, {"x", "x", "o"}, {"x", "o", "o"}},
			expectedMoveWeight: -1,
		},
		{
			description:       "my win by diag",
			board:             [][]string{{"x", "x", "o"}, {"x", "x", "o"}, {"o", "o", "x"}},
			expectedMoveWeight: 1,
		},



	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			computerSign ="x"
			resultWeight := weightMove(c.board)


			if c.expectedMoveWeight != resultWeight {
				t.Errorf("\nUnexpected resultRow: \n\t\t expected: %d, \n\t\t actual: %d", c.expectedMoveWeight, resultWeight)

			}
		})
	}
}

func TestIsRemaining(t *testing.T) {

	cases := []struct {
		description       string
		board             [][]string
		isRemaining			bool
	}{
		{
			description:       "no moves remaining when all are filled",
			board:             [][]string{{"x", "x", "o", "x"}, {"o", "x", "x", "o"}, {"x", "x", "o", "o"}, {"x", "x", "o", "o"}},
			isRemaining: 	   false,
		},
		{
			description:       "remaining moves when multiple are free",
			board:             [][]string{{"x", "_", "x"}, {"_", "_", "_"}, {"o", "_", "o"}},
			isRemaining: 	   true,
		},
		{
			description:       "remaining moves when all are filled except 1",
			board:             [][]string{{"x", "x", "o"}, {"o", "o", "o"}, {"x", "x", "_"}},
			isRemaining: 	   true,
		},
		{
			description:       "remaining moves when complete board is empty",
			board:             [][]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}},
			isRemaining: 	   true,
		},



	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			computerSign ="x"
			isRemaining := isRemaining(c.board)


			if c.isRemaining != isRemaining {
				t.Errorf("\nUnexpected resultRow: \n\t\t expected: %d, \n\t\t actual: %d", c.isRemaining, isRemaining)

			}
		})
	}
}

func TestMinMax(t *testing.T) {

	cases := []struct {
		description       string
		a             	int
		b 				int
		min				int
		max 			int
	}{
		{
			description:       "return max and min of two numbers",
			a:             	45,
			b:				23,
			min: 	   23,
			max:  		45,
		},
	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			computerSign ="x"
			max := Max(c.a,c.b)
			min := Min(c.a,c.b)


			if c.max != max {
				t.Errorf("\nUnexpected max: \n\t\t expected: %d, \n\t\t actual: %d", c.max, max)
			}
			if c.min != min {
				t.Errorf("\nUnexpected min: \n\t\t expected: %d, \n\t\t actual: %d", c.min, min)
			}
		})
	}
}

func TestReadArgs(t *testing.T) {

	cases := []struct {
		description       string
		boardSize         int
		expectedBoardSize int
		player1Sign       string
		player2Sign       string
		computerSign      string

	}{
		{
			description:       "Should return given size as board size without error",
			expectedBoardSize:5,
		},
	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {

			readArgs("test_config.txt")



			if c.boardSize != boardSize {
				t.Errorf("\nUnexpected board size: \n\t\t expected: %d, \n\t\t actual: %d", c.boardSize, boardSize)
			}
			if c.computerSign != computerSign || c.player1Sign != opponent1Sign || c.player2Sign != opponent2Sign {
				t.Errorf("\nUnexpected board signs: \n\t\t expected: %d, \n\t\t actual: %d \n\t\t expected: %d, \n\t\t actual: %d \n\t\t expected: %d, \n\t\t actual: %d", c.computerSign, computerSign, c.player1Sign, opponent1Sign, c.player2Sign, opponent2Sign)
			}
		})
	}
}

func TestFinder(t *testing.T) {

	cases := []struct {
		description       string
		board             [][]string
		expectedResult int
		isMyTurn	bool
	}{
		{
			description:       "should return opponents win for the given board",
			board:             [][]string{	{"x", "x", "o"},
											{"o", "o", "o"},
											{"x", "x", "_"}},

			expectedResult:-1,
			isMyTurn:false,
		},
		{
			description:       "should return my win for the given diag board",
			board:             [][]string{	{"x", "o", "x"},
											{"o", "x", "o"},
											{"x", "o", "_"}},
			expectedResult:1,
			isMyTurn:true,
		},
		{
			description:       "should return my win for the given board",
			board:             [][]string{	{"x", "x", "o"},
											{"x", "o", "o"},
											{"x", "o", "_"}},
			expectedResult:1,
			isMyTurn:true,
		},
		{
			description:       "should return oponnents win for the given row board",
			board:             [][]string{	{"x", "x", "o"},
											{"o", "o", "o"},
											{"x", "o", "_"}},
			expectedResult:-1,
			isMyTurn:false,
		},
		{
			description:       "should return oponnents win for the given diag board",
			board:             [][]string{	{"o", "x", "o"},
											{"o", "o", "x"},
											{"x", "o", "o"}},
			expectedResult:-1,
			isMyTurn:false,
		},
		{
			description:       "should return oponnents win for the given diag part empty board",
			board:             [][]string{	{"o", "x", "_"},
											{"x", "o", "x"},
											{"_", "_", "o"}},
			expectedResult:-1,
			isMyTurn:false,
		},
		{
			description:       "should return my win for the given row part empty board",
			board:             [][]string{	{"o", "x", "_"},
											{"_", "x", "o"},
											{"_", "x", "o"}},
			expectedResult:1,
			isMyTurn:false,
		},
		{
			description:       "should return draw for the given non empty board",
			board:             [][]string{	{"o", "x", "x"},
											{"x", "o", "o"},
											{"x", "o", "x"}},
			expectedResult:0,
			isMyTurn:false,
		},
		{
			description:       "should return my possible win (row-column) when its my turn",
			board:             [][]string{	{"_", "x", "x"},
											{"x", "o", "o"},
											{"x", "o", "x"}},
			expectedResult:1,
			isMyTurn:true,
		},
		{
			description:       "should return my possible win (column) when its my turn",
			board:             [][]string{	{"_", "o", "x"},
											{"x", "o", "o"},
											{"x", "x", "o"}},
			expectedResult:1,
			isMyTurn:true,
		},
		{
			description:       "should return opponents win (column) when its opponents turn",
			board:             [][]string{	{"_", "o", "x"},
											{"x", "o", "o"},
											{"x", "x", "o"}},
			expectedResult:-1,
			isMyTurn:false,
		},
		{
			description:       "should return opponents win (column) when its opponents turn",
			board:             [][]string{	{"o", "o", "x"},
											{"x", "o", "o"},
											{"x", "x", "_"}},
			expectedResult:-1,
			isMyTurn:false,
		},

		/*{
			description:       "should return opponents win (column) when its opponents turn",
			board:             [][]string{	{"_", "_", "_"},
											{"_", "o", "_"},
											{"_", "_", "_"},},
			expectedResult:-1,
			isMyTurn:false,
		},*/


	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			computerSign ="x"
			opponent1Sign="o"
			result:=finder(c.board,c.isMyTurn,2)

			if c.expectedResult != result {
				t.Errorf("\nUnexpected win: \n\t\t expected: %d, \n\t\t actual: %d", c.expectedResult, result)
			}
		})
	}
}
func TestPlay(t *testing.T) {

	cases := []struct {
		description       string
		board             [][]string
		expectedRow int
		expectedCol	int
	}{
		{
			description:       "should return the winning move for computer",
			board:             [][]string{	{"x", "x", "o"},
											{"o", "_", "_"},
											{"o", "x", "_"}},

			expectedRow:1,
			expectedCol:1,
		},
		{
			description:       "should return the winning move for computer",
			board:             [][]string{	{"x", "x", "_"},
											{"o", "o", "_"},
											{"o", "x", "_"}},

			expectedRow:0,
			expectedCol:2,
		},
		{
			description:       "should return the move not letting opponent win",
			board:             [][]string{	{"x", "_", "_"},
											{"o", "o", "_"},
											{"_", "_", "x"}},

			expectedRow:1,
			expectedCol:2,
		},
		{
			description:       "should return a move for the given board",
			board:             [][]string{	{"_", "_","_"},
											{"_", "o","_"},
											{"_", "_","_"}},

			expectedRow:0,
			expectedCol:0,
		},

	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			computerSign ="x"
			opponent1Sign="o"
			expectedRow, expectedCol:=play(c.board)

			printBoard(c.board)
			if c.expectedCol != expectedCol {
				t.Errorf("\nUnexpected column: \n\t\t expected: %d, \n\t\t actual: %d", c.expectedCol, expectedCol)
			}
			if c.expectedRow != expectedRow {
				t.Errorf("\nUnexpected row: \n\t\t expected: %d, \n\t\t actual: %d", c.expectedRow, expectedRow)
			}
		})
	}
}