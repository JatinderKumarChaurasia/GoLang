package main

import (
	"fmt"
	"math/rand"
	"time"
)

type boardTicTac [3][3]rune

func main() {
	rand.Seed(time.Now().UnixNano())
	var playerMove bool
	var playerWon string
	var winner bool

	var board boardTicTac

	fmt.Println("Starting Game : Board Empty ")
	board.displayBoard()

	if rand.Intn(2) == 0 {
		playerMove = true
	} else {
		playerMove = false
	}

	for i := 0; i < 9; i++ {
		if playerMove {
			fmt.Println("Player Move : ", i+1)
			board.player()
			playerMove = false
		} else {
			fmt.Println("Computer Move : ", i+1)
			time.Sleep(time.Second)
			board.computer()
			playerMove = true
		}
		if playerWon, winner = board.check(); winner {
			break
		}
		board.displayBoard()
	}

	fmt.Printf("******%v Won******\nFinal Board View : \n", playerWon)
}

func (board boardTicTac) displayBoard() {
	fmt.Print("------------------")
	for i := 0; i < 3; i++ {
		fmt.Print("\n| ")
		for j := 0; j < 3; j++ {
			fmt.Printf(" %c | ", board[i][j])
		}
		fmt.Print("\n------------------")
	}
	fmt.Println("")
}

func (board *boardTicTac) player() {
	var x, y int
	fmt.Println("Enter the Row[1-3] and Column[1-3] Positions : ")
	if _, err := fmt.Scan(&x, &y); err == nil {
		x--
		y--
		if (x >= 0 && x <= 2) && (y >= 0 && y <= 2) && (board[x][y] == 0) {
			board[x][y] = 'X'
		} else {
			fmt.Println("Invalid Player Input or Position Not Empty. Try Again !")
			board.player()
		}
	} else {
		fmt.Println("Invalid Player Input or Position Not Empty. Try Again !")
		board.player()
	}
}

func (board *boardTicTac) computer() {
	var x, y int
	for {
		x = rand.Intn(3)
		y = rand.Intn(3)
		if board[x][y] == 0 {
			board[x][y] = 'O'
			break
		}
	}
}

func (board *boardTicTac) check() (string, bool) {
	for i := 0; i < 3; i++ {
		if (rune(board[i][0]) == 'X') && (board[i][0] == board[i][1]) && (board[i][0] == board[i][2]) {
			return "Player", true
		} else if (rune(board[i][0]) == 'O') && (board[i][0] == board[i][1]) && (board[i][0] == board[i][2]) {
			return "Computer", true
		}
	}
	for i := 0; i < 3; i++ {
		if (rune(board[0][i]) == 'X') && (board[0][i] == board[1][i]) && (board[0][i] == board[2][i]) {
			return "Player", true
		} else if (rune(board[0][i]) == 'O') && (board[0][i] == board[1][i]) && (board[0][i] == board[2][i]) {
			return "Computer", true
		}
	}
	if ((rune(board[0][0]) == 'X') && (board[0][0] == board[1][1]) && (board[1][1] == board[2][2])) || ((rune(board[0][2]) == 'X') && (board[0][2] == board[1][1]) && (board[1][1] == board[2][0])) {
		return "Player", true
	} else if ((rune(board[0][0]) == 'O') && (board[0][0] == board[1][1]) && (board[1][1] == board[2][2])) || ((rune(board[0][2]) == 'O') && (board[0][2] == board[1][1]) && (board[1][1] == board[2][0])) {
		return "Computer", true
	}
	return "Draw", false
}
