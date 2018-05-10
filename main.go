package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"learn/tic/players"
)

/*var computerSign string
var opponent1Sign string
var opponent2Sign string*/

var boardSize int
//var players []Player

func main() {

	//players:=make([]Player,0)
	boardSize = readArgs("config.txt")


	board := initializeBoard(boardSize, boardSize)

	//printBoard(board)

	// Can be used to change number of players
	/*for {
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

	}*/
	for {
	for i:= range players.GetPlayers() {
		players.GetPlayers()[i].Move(board)

	}}

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
			//opponent1Sign = propertyVal
			players.SetPlayers(players.ManualPlayer{propertyVal})
			//players=append(getPlayers(),manualPlayer{propertyVal})
		case "Player 2 Char":
			//opponent2Sign = propertyVal
			players.SetPlayers(players.ManualPlayer{propertyVal})
			//players=append(players,manualPlayer{propertyVal})
		case "Computer Char":
			//computerSign = propertyVal
			players.SetPlayers(players.ComputerPlayer{propertyVal})
			//players=append(players,computerPlayer{propertyVal})

		}

	}

	if boardSize == -1 {
		os.Exit(1)
	}

	fmt.Printf("board size : %d ", boardSize)
	fmt.Println()
	fmt.Println("Let the game begin !!")

	return boardSize

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
	players.PrintBoard(board)
	return board
}




