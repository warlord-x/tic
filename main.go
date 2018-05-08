package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var computerSign string
var opponent1Sign string
var opponent2Sign string

var boardSize int

func main() {

	boardSize = readArgs("config.txt")
	board := initializeBoard(boardSize, boardSize)

	printBoard(board)

	// Can be used to change number of players
	for {
		fmt.Println("ready player 1..")
		readAndMarkMove(board, opponent1Sign)
		fmt.Println("Your move")
		printBoard(board)
		fmt.Println("ready player 2..")
		readAndMarkMove(board, opponent2Sign)
		fmt.Println("player 2's  move")
		printBoard(board)
		fmt.Println("My move")
		row, col := play(board)
		board[row][col] = computerSign
		printBoard(board)
		checkWin(board, computerSign)

	}

}

func readPrint(b [][]string) [2]int {

	input := [2]int{-1, -1}

	for input[0] == -1 || input[1] == -1 {
		reader := bufio.NewReader(os.Stdin)
		var storageString string
		storageString, _ = reader.ReadString('\n')
		test := strings.Split(storageString, ",")

		if len(test) > 2 {
			fmt.Println("incorrect input..try again in format number,number")
			input[0] = -1
			input[1] = -1
			continue
		}
		for a := range test {
			val := strings.TrimSpace(test[a])
			convtered, err := strconv.Atoi(val)
			if err == nil && convtered < boardSize && convtered >= 0 {
				input[a] = convtered
			} else {
				fmt.Println("incorrect input..try again")
				input[0] = -1
				input[1] = -1
				break
			}
		}
		if !(input[0] == -1 && input[1] == -1) {
			if b[input[0]][input[1]] != "_" {
				fmt.Println("incorrect input , value already exist")
				input[0] = -1
				input[1] = -1
			}
		}

	}
	return input
}
func readArgs(filename string) int {

	b, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	s := string(b)
	lines := strings.Split(s, "\n")

	for line := range lines {

		readLine := strings.TrimSpace(lines[line])
		keyVal := strings.Split(readLine, ":")

		propertyName := strings.TrimSpace(keyVal[0])

		propertyVal := strings.TrimSpace(strings.TrimSpace(keyVal[1]))

		switch propertyName {
		case "Board Size":
			{
				readSize, err := strconv.Atoi(strings.TrimSpace(propertyVal))
				if err != nil || readSize < 3 || readSize > 10 {

					fmt.Println("incorrect input config file")
					os.Exit(1)
				} else {
					boardSize = readSize
				}
			}
		case "Player 1 Char":
			opponent1Sign = propertyVal
		case "Player 2 Char":
			opponent2Sign = propertyVal
		case "Computer Char":
			computerSign = propertyVal

		}

	}

	if boardSize == -1 {
		os.Exit(1)
	}

	fmt.Printf("board size : %d , computer char %s , player1 char %s , player 2 char %s", boardSize, computerSign, opponent1Sign, opponent2Sign)
	fmt.Println()
	fmt.Println("Let the game begin !!")

	return boardSize

}

func readAndMarkMove(b [][]string, opponentSign string) {

	input := readPrint(b)

	b[input[0]][input[1]] = opponentSign

	checkWin(b, opponentSign)

}

func checkWin(b [][]string, sign string) {
	switch weightMove(b) {
	case 1:
		fmt.Println("Game over. I win!!")
		os.Exit(1)
	case -1:

		if sign == opponent1Sign {
			fmt.Println("Game over. Player 1 wins!!")
		} else if sign == opponent2Sign {
			fmt.Println("Game over. Player 2 wins!!")
		}
		os.Exit(1)
	default:
		if !isRemaining(b) {
			fmt.Println("no moves remaining ..exiting")
			os.Exit(1)
		}
	}

}
func initializeBoard(r, c int) [][]string {

	board := make([][]string, 0)
	for i := 0; i < r; i++ {
		row := make([]string, 0)
		for j := 0; j < c; j++ {
			row = append(row, "_")
		}
		board = append(board, row)
	}

	return board
}

func printBoard(b [][]string) {
	for i := range b {
		fmt.Println(b[i])
	}
}

func isRemaining(b [][]string) bool {

	for i := range b {
		for j := range b[i] {
			if b[i][j] == "_" {
				return true
			}
		}
	}
	return false
}

func play(b [][]string) (int, int) {

	rowX := -1
	columnX := -1

	highVal := -10000

	for row := range b {
		for column := range b[0] {
			if b[row][column] == "_" {
				b[row][column] = computerSign
				val := finder(b, false, 1)
				b[row][column] = "_"
				if val > highVal {
					highVal = val
					rowX = row
					columnX = column
				}
			}
		}
	}
	return rowX, columnX

}

func finder(b [][]string, isMe bool, recursionLevel int) int {
	score := weightMove(b)

	if score == 1 || score == -1 || recursionLevel == 0 {
		return score
	}
	if !isRemaining(b) {
		return 0
	}

	if isMe {
		best := -1000
		for row := range b {
			for column := range b[0] {
				if b[row][column] == "_" {
					b[row][column] = computerSign
					best = Max(finder(b, !isMe, recursionLevel-1), best)
					b[row][column] = "_"
				}
			}
		}
		return best
	} else {
		best := 1000
		for row := range b {
			for column := range b[0] {
				if b[row][column] == "_" {
					b[row][column] = opponent1Sign
					best = Min(finder(b, !isMe, recursionLevel-1), best)
					b[row][column] = "_"
				}
			}
		}

		for row := range b {
			for column := range b[0] {
				if b[row][column] == "_" {
					b[row][column] = opponent2Sign
					best = Min(finder(b, !isMe, recursionLevel-1), best)
					b[row][column] = "_"
				}
			}
		}

		return best
	}
}

func weightMove(b [][]string) int {

	V := checkRow(b)
	if V == 0 {
		V = checkColumn(b)
		if V == 0 {
			return checkDiag(b)
		}
	}
	return V
}

func checkRow(b [][]string) int {
	result := 0
	for row := range b {
		notMatched := true
		k := len(b) - 1 //TODO : assumes symetric board
		i := 0
		for ; i < k && notMatched; i++ {

			if b[row][i] != b[row][i+1] {
				notMatched = false
			}

		}

		if notMatched {

			if b[row][i] == computerSign {
				return 1
			} else if b[row][i] == opponent1Sign || b[row][i] == opponent2Sign {
				return -1
			}
		}

	}
	return result
}
func checkColumn(b [][]string) int {
	result := 0
	for column := range b[0] {
		notMatched := true
		k := len(b[0]) - 1
		i := 0
		for ; i < k && notMatched; i++ {

			if b[i][column] != b[i+1][column] {
				notMatched = false
			}

		}

		if notMatched {
			if b[i][column] == computerSign {
				return 1
			} else if b[i][column] == opponent1Sign || b[i][column] == opponent2Sign {
				return -1
			}
		}

	}
	return result
}

func checkDiag(b [][]string) int {
	result := 0
	notMatched := true
	k := len(b) - 1
	i := 0
	for ; i < k && notMatched; i++ {

		if b[i][i] != b[i+1][i+1] {
			notMatched = false
		}

	}
	if notMatched {
		if b[i][i] == computerSign {
			return 1
		} else if b[i][i] == opponent1Sign || b[i][i] == opponent2Sign {
			return -1
		}
	}
	if result == 0 {
		j := k
		notMatched = true
		i := 0
		for ; i < k && notMatched; i++ {

			if b[j][i] != b[j-1][i+1] {
				notMatched = false
			}
			j--
		}
		if notMatched {
			if b[j][i] == computerSign {
				return 1
			} else if b[j][i] == opponent1Sign || b[j][i] == opponent2Sign {
				return -1
			}
		}
	}
	return result
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
