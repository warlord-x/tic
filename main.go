package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ball struct {
	count int
}

type BoardSpecs struct {
	row    int
	column int
}

var mySign string
var opponent1Sign string

func main() {

	mySign = "x"
	opponent1Sign = "o"
	board := createDummyBoard()
	printBoard(board)
	notDone:=true
	for isRemaining(board) && notDone {
		readAndMarkMove(board)
		fmt.Println("Your move")
		printBoard(board)
		if isRemaining(board) {
			fmt.Println("My move")
			row,col := play(board)
			board[row][col]=mySign
			printBoard(board)
		}
		s:=weightMove(board)
		fmt.Println("score::::::::::",s)
		if s==1 || s==-1 {
			notDone=false
		}
	}

}

func readPrint(b [][]string) [2]int {

	input:=[2]int{-1,-1}


	for input[0]==-1 || input[1]==-1 {
		reader := bufio.NewReader(os.Stdin)
		var storageString string
		storageString, _ = reader.ReadString('\n')
		fmt.Println("readvalue:", storageString)
		test := strings.Split(storageString, ",")

		if len(test) > 2 {
			fmt.Println("incorrect input..try")
			input[0]=-1
			input[1]=-1
			continue
		}
		for a := range test {
			val := strings.TrimSpace(test[a])
			convtered, err := strconv.Atoi(val)
			if err == nil {
				fmt.Printf("%d looks like a number.\n", convtered)
				input[a]=convtered
				fmt.Println("lentth of input", len(input))
			} else {
				fmt.Println("incorrect input try again")
				input[0]=-1
				input[1]=-1
				break
			}
		}
		if !(input[0]==-1 && input[1]==-1) {
			if b[input[0]][input[1]]!="_" {
				fmt.Println("incorrect input , value already exist")
				input[0]=-1
				input[1]=-1
			}
		}

	}
	return input
}

func readAndMarkMove(b [][]string) {
	/*row := -1
	col := -1
	_, err := fmt.Scanf("%d", &row)
	if nil != err {
		panic(err)
	}
	fmt.Scanf("%d", &col)*/

	//if b[row][col] != "_"
	input:=readPrint(b)

	b[input[0]][input[1]] = "o"

}
func player(action string, table chan *Ball) {
	for {
		ball := <-table
		ball.count++
		fmt.Println(action, ball.count)
		time.Sleep(time.Millisecond * 100)
		table <- ball
	}
}

func initailize(r, c int, mysign string, oppo1sign string) {

	mySign = mysign
	opponent1Sign = oppo1sign

	row := make([]string, 0)
	for i := 0; i < r; i++ {
		row = append(row, "_")
	}

	col := make([][]string, 0)

	for j := 0; j < c; j++ {
		col = append(col, row)
	}

	printBoard(col)
}

func printBoard(b [][]string) {
	for i := range b {
		fmt.Println(b[i])
	}
	fmt.Println("++++++++++++++++")
}

func createDummyBoard() [][]string {
	//a := [][]string{{"o", "o", "o", "o"}, {"k", "j", "e", "b"}, {"k", "j", "e", "b"}, {"m", "d", "b", "b"}}
	//Column
	//a := [][]string{{"o", "x", "x", "o"},{"x", "x", "o", "x"},{"o", "x", "x", "x"},{"x", "x", "x", "x"}}
	//Diag
	//a := [][]string{{"o", "x", "x", "o"}, {"x", "o", "o", "x"}, {"o", "x", "o", "x"}, {"x", "x", "x", "o"}}
	//Reverse Diag
	//a := [][]string{{"o", "x", "x", "o"}, {"x", "x", "o", "x"}, {"o", "o", "o", "x"}, {"o", "x", "x", "o"}}
	//Dummy Empty
	//a := [][]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}
	//Dummy Empty 4
	a := [][]string{{"_", "_", "_","_"}, {"_","_", "_", "_"}, {"_","_", "_", "_"},{"_","_", "_", "_"}}

	return a

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
		//if row==0 {
		notMatched := true
		k := len(b) - 1 //TODO : assumes symetric board
		i := 0
		for ; i < k && notMatched; i++ {
			//fmt.Println("first:", b[row][i], "second:", b[row][i+1])

			if b[row][i] != b[row][i+1] {
				notMatched = false
			}

		}

		if notMatched {

			if b[row][i] == "x" {
				return 1
			} else if b[row][i] == "o" {
				return -1
			}
		}

	}
	//fmt.Println(result)
	return result
}
func checkColumn(b [][]string) int {
	//printBoard(b)
	result := 0
	for column := range b[0] {
		notMatched := true
		k := len(b[0]) - 1 //TODO : assumes symetic board
		i := 0
		for ; i < k && notMatched; i++ {
			//fmt.Println("first:", b[i][column], "second:", b[i+1][column])

			if b[i][column] != b[i+1][column] {
				notMatched = false
			}

		}

		if notMatched {
			if b[i][column] == "x" {
				return 1
			} else if b[i][column] == "o" {
				return -1
			}
		}

	}
	//fmt.Println(result)
	return result
}

func checkDiag(b [][]string) int {
	result := 0
	notMatched := true
	k := len(b) - 1
	i := 0
	for ; i < k && notMatched; i++ {

		//fmt.Println("first:", b[i][i], "second:", b[i+1][i+1])

		if b[i][i] != b[i+1][i+1] {
			notMatched = false
		}

	}
	if notMatched {
		if b[i][i] == "x" {
			return 1
		} else if b[i][i] == "o" {
			return -1
		}
	}
	if result == 0 {
		j := k
		notMatched = true
		i := 0
		for ; i < k && notMatched; i++ {

			//fmt.Println("first:", b[j][i], "second:", b[j-1][i+1])

			if b[j][i] != b[j-1][i+1] {
				notMatched = false
			}

			j--
		}
		if notMatched {
			if b[j][i] == "x" {
				return 1
			} else if b[j][i] == "o" {
				return -1
			}
		}
	}
	//fmt.Println(result)
	return result
}

func play(b [][]string) (int, int) {

	rowX := -1
	columnX := -1

	highVal := -10000

	for row := range b {
		for column := range b[0] {
			if b[row][column] == "_" {
				b[row][column] = mySign
				//printBoard(b)
				val := finder(b, false,1)
				b[row][column] = "_"
				if val > highVal {
					highVal = val
					rowX = row
					columnX = column
				}
			}
		}
	}
fmt.Println("printing before exiting play",rowX,columnX)
	printBoard(b)
	//b[rowX][columnX] = mySign
	return rowX, columnX

}

func finder(b [][]string, isMe bool,recursionLevel int) int {
	score := weightMove(b)

	if score == 1 || score == -1 || recursionLevel == 0 {
		//fmt.Println("returning score:", score)
		return score
	}
	if !isRemaining(b) {
		//fmt.Println("no moves remaining")
		return 0
	}

	if isMe {
		best := -1000
		for row := range b {
			for column := range b[0] {
				if b[row][column] == "_" {
					b[row][column] = mySign
					//printBoard(b)
					best = Max(finder(b, !isMe,recursionLevel-1), best)
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
					b[row][column] = opponent1Sign //TODO : other signs
					//printBoard(b)
					best = Min(finder(b, !isMe,recursionLevel-1), best)
					b[row][column] = "_"
				}
			}
		}
		return best
	}
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
