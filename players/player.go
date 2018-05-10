package players

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

var players []Player

type Player interface {
	Move([][]string)
	GetSign() string
	isComputer() bool
}

type ManualPlayer struct {
	Sign string
}

type ComputerPlayer struct {
	Sign string
}

func (p ManualPlayer) Move(b [][]string){
	fmt.Println("manual player moved with Sign :",p.Sign)
	readAndMarkMove(b,p.Sign)
}

func (p ComputerPlayer) Move(b [][]string){
	fmt.Println("computer player moved with Sign :",p.Sign)

	play(b,p)
	checkWin(b, p.Sign)
}

func (p ManualPlayer) GetSign() string {
	return p.Sign
}
func (p ComputerPlayer) GetSign() string {
	return p.Sign
}
func (p ComputerPlayer) isComputer() bool {
	return true
}
func (p ManualPlayer) isComputer() bool {
	return false
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
			if err == nil && convtered < len(b) && convtered >= 0 {
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

/*func Move(p player){
	p.Move()
}*/

func readAndMarkMove(b [][]string, opponentSign string) {

	input := readPrint(b)

	b[input[0]][input[1]] = opponentSign

	checkWin(b, opponentSign)

}

func play(b [][]string,p Player) (int, int) {

	rowX := -1
	columnX := -1

	highVal := -10000

	for row := range b {
		for column := range b[0] {
			if b[row][column] == "_" {
				b[row][column] = getComputerSign()
				val := finder(b, false, 1,p)
				b[row][column] = "_"
				if val > highVal {
					highVal = val
					rowX = row
					columnX = column
				}
			}
		}
	}
	b[rowX][columnX] = p.GetSign()
	return rowX, columnX

}


func finder(b [][]string, isMe bool, recursionLevel int, p Player) int {
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

					b[row][column] = getComputerSign()
					best = Max(finder(b, !isMe, recursionLevel-1,p), best)
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

					for i:= range GetPlayers() {
						if !players[i].isComputer(){
							b[row][column] = players[i].GetSign()
							best = Min(finder(b, !isMe, recursionLevel-1,p), best)
							b[row][column] = "_"
						}
					}

					/*b[row][column] = opponent1Sign
					best = Min(finder(b, !isMe, recursionLevel-1), best)
					b[row][column] = "_"
					b[row][column] = opponent2Sign
					best = Min(finder(b, !isMe, recursionLevel-1), best)
					b[row][column] = "_"*/
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



			if b[row][i] == getComputerSign() {
				return 1
			} else if inOtherPlayersSign(b[row][i]){
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
			if b[i][column] == getComputerSign() {
				return 1
			} else if inOtherPlayersSign(b[i][column]){
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
		if b[i][i] == getComputerSign() {
			return 1
		} else if inOtherPlayersSign(b[i][i]){
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
			if b[j][i] == getComputerSign() {
				return 1
			} else if inOtherPlayersSign(b[j][i]){
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
func checkWin(b [][]string, sign string) {
	PrintBoard(b)
	switch weightMove(b) {
	case 1:
		fmt.Println("Game over. I win!!")

		os.Exit(1)
	case -1:

		//if Sign == opponent1Sign {
			fmt.Printf("Game over. Player with Sign %s wins!!",sign)
		/*} else if Sign == opponent2Sign {
			fmt.Println("Game over. Player 2 wins!!")
		}*/
		os.Exit(1)
	default:
		if !isRemaining(b) {
			fmt.Println("no moves remaining ..exiting")
			os.Exit(1)
		}
	}

}


func GetPlayers() []Player {
	return players
}

func getComputerSign() string {
	var computerSign string
	for i:= range players {
		if players[i].isComputer() {
			computerSign = players[i].GetSign()
		}
	}
	return computerSign
}

func inOtherPlayersSign(sign string) bool {
	found:= false
	for i:= range players {
		if ! players[i].isComputer() {
			if sign == players[i].GetSign() {
				found = true
			}
		}
	}
	return found
}

func SetPlayers(p Player){
	if nil == players {
		players = make([]Player,0)
	}
	players=append(players,p)
}

func PrintBoard(b [][]string) {
	for i := range b {
		fmt.Println(b[i])
	}
}