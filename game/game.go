package game

import (
	"log"
	"fmt"
	"strconv"
)

type GameInfo struct {
	Wins int
	Losses int
	Ties int
	Board [9]string
}

type Player struct {
	Name string
	Password string
	Ingame bool
	GameData GameInfo
}

var Players map[string]*Player

var playerLetter = "X"
var computerLetter = "O"

var winners = [][3]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

func init() {
	Players = make(map[string]*Player, 0)
	p := &Player{Name: "paul", Password: "test123", Ingame: false, GameData: GameInfo{Wins: 0, Losses: 0, Ties: 0 } }
	Players["paul"] = p
	p = &Player{Name: "carol", Password: "test123", Ingame: false, GameData: GameInfo{Wins: 0, Losses: 0, Ties: 0 } }
	Players["carol"] = p

}

func NewPlayer(username string, password string) (*Player, string) {
	if pl, ok := Players[username]; ok {
		return pl, "PlayerExists"
	}
	p := &Player{Name: username, Password: password, Ingame: false, GameData: GameInfo{Wins: 0, Losses: 0, Ties: 0 }}
	Players[username] = p
	return Players[username], ""
}

func StartGame(username string) {
	Players[username].Ingame = true
}

func PlayerMove(square string, username string) {
	log.Println("Move:", square)
	i, err := strconv.Atoi(square)
	i--
	if err != nil {
		fmt.Println(err)
	}
	if!Players[username].Ingame {
		fmt.Printf("You are not in a game\n", square)
	} else {
		fmt.Printf("You played square: %s\n", i)
		Players[username].GameData.Board[i] = playerLetter
		fmt.Printf("Board: %v\n", Players[username].GameData.Board)
		if IsWinner(Players[username].GameData.Board, playerLetter) {
			Players[username].Ingame = false
			Players[username].GameData.Wins++
			Players[username].GameData.Board = [9]string{}
			fmt.Println("Player Won!!")
		} else if IsBoardFull(Players[username].GameData.Board) {
			Players[username].Ingame = false
			Players[username].GameData.Ties++
			Players[username].GameData.Board = [9]string{}
			fmt.Println("It's a tie")
		} else {
			computerMove(username)
			if IsWinner(Players[username].GameData.Board, computerLetter) {
				Players[username].Ingame = false
				Players[username].GameData.Losses++
				Players[username].GameData.Board = [9]string{}
				fmt.Println("Computer Won!!")
			}else if IsBoardFull(Players[username].GameData.Board) {
				Players[username].Ingame = false
				Players[username].GameData.Ties++
				Players[username].GameData.Board = [9]string{}
				fmt.Println("It's a tie")
			}
		}
	}
}

func computerMove(username string) {
	fmt.Printf("Computer plays now!!\n")

	//If one move away (play or block)
	i := IsOneMoveAway(Players[username].GameData.Board, computerLetter)
	log.Println("1 - i: ", i)
	if i == -1 {
		i = IsOneMoveAway(Players[username].GameData.Board, playerLetter)
		log.Println("2 - i: ", i)
	}
	if i == -1 {
		i = ChooseRandomMoveFromList(Players[username].GameData.Board, []int{0, 2, 6, 8})
		log.Println("3 - i: ", i)
	}
	if i == -1 {
		i = ChooseRandomMoveFromList(Players[username].GameData.Board, []int{4})
		log.Println("4 - i: ", i)
	}
	if i == -1 {
		i = ChooseRandomMoveFromList(Players[username].GameData.Board, []int{1, 3, 5, 7})
		log.Println("5 - i: ", i)
	}
	if i != -1 {
		Players[username].GameData.Board[i] = "O"
	} else {
		log.Println("No move available")
	}
}

func ChooseRandomMoveFromList(bo [9]string, choices []int) int {
	for x := range choices {
		log.Printf("choose: %d %d |%s|", x, choices[x], bo[choices[x]])
		if bo[choices[x]] == "" {
			return choices[x]
		}
	}
	return -1
}


func IsOneMoveAway(bo [9]string, le string) int {
	fmt.Printf("Check - %v\n", bo)
	if bo[0] == "" && bo[1] == le && bo[2] == le {
		return 0
	} else if bo[0] == le && bo[1] == "" && bo[2] == le {
		return 1
	} else if bo[0] == le && bo[1] == le && bo[2] == "" {
		return 2
	} else if bo[3] == "" && bo[4] == le && bo[5] == le {
		return 3
	} else if bo[3] == le && bo[4] == "" && bo[5] == le {
		return 4
	} else if bo[3] == le && bo[4] == le && bo[5] == "" {
		return 5
	} else if bo[6] == "" && bo[7] == le && bo[8] == le {
		return 6
	} else if bo[6] == le && bo[7] == "" && bo[8] == le {
		return 7
	} else if bo[6] == le && bo[7] == le && bo[8] == "" {
		return 8
	} else if bo[0] == "" && bo[3] == le && bo[6] == le {
		return 0
	} else if bo[0] == le && bo[3] == "" && bo[6] == le {
		return 3
	} else if bo[0] == le && bo[3] == le && bo[6] == "" {
		return 6
	} else if bo[1] == "" && bo[4] == le && bo[7] == le {
		return 1
	} else if bo[1] == le && bo[4] == "" && bo[7] == le {
		return 4
	} else if bo[1] == le && bo[4] == le && bo[7] == "" {
		return 7
	} else if bo[2] == "" && bo[5] == le && bo[8] == le {
		return 2
	} else if bo[2] == le && bo[5] == "" && bo[8] == le {
		return 5
	} else if bo[2] == le && bo[5] == le && bo[8] == "" {
		return 8
	} else if bo[0] == "" && bo[4] == le && bo[8] == le {
		return 0
	} else if bo[0] == le && bo[4] == "" && bo[8] == le {
		return 4
	} else if bo[0] == le && bo[4] == le && bo[8] == "" {
		return 8
	} else if bo[2] == "" && bo[4] == le && bo[6] == le {
		return 2
	} else if bo[2] == le && bo[4] == "" && bo[6] == le {
		return 4
	} else if bo[2] == le && bo[4] == le && bo[6] == "" {
		return 6
	} else {
		return -1
	}
}


func IsWinner(bo [9]string, le string) bool {
	for _, pos := range winners {
		if bo[pos[0]] == le && bo[pos[1]] == le && bo[pos[2]] == le {
			return true
		}
	}
	return false
}

func IsBoardFull(bo [9]string) bool {
	for i := 0; i < 9; i++ {
		if bo[i] == "" {
			return false
		}
	}
	return true
}

