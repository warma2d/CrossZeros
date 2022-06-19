package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var PlayerSymbols = [2]string{"X", "0"}

func main() {
	var player1Name = userInput("Enter name 1st player: ")
	var player2Name = userInput("Enter name 2nd player: ")

	playerNames := [2]string{player1Name, player2Name}
	table := [3][3]string{}

	fillTildaTable(&table)
	fmt.Println("")
	drawTable(table)
	play(playerNames, table)
}

func play(playerNames [2]string, table [3][3]string) {
	var playing = true
	var activePlayerNumber = 0

	for playing {
		printTurnPlayer(playerNames[activePlayerNumber])

	turn:
		{
			i, j, err := getIndexesFromPlayer()
			if err != nil {
				fmt.Println(err.Error())
				goto turn
			}

			err = doTurn(&table, activePlayerNumber, i, j)
			if err != nil {
				fmt.Println(err.Error())
				goto turn
			}
		}

		changePlayer(&activePlayerNumber)

		if isGameOver(table) {
			playing = false
		}

		drawTable(table)
	}

	fmt.Println("Game over!")
	userInput("")
}

func changePlayer(activePlayerNumber *int) {
	if *activePlayerNumber == 0 {
		*activePlayerNumber = 1
	} else {
		*activePlayerNumber = 0
	}
}

func isGameOver(table [3][3]string) bool {
	if isWonSymbol(PlayerSymbols[0], table) || isWonSymbol(PlayerSymbols[1], table) {
		return true
	}

	usedFields := 0
	maxUsedFields := 9

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if table[i][j] != "~" {
				usedFields++
			}
		}
	}

	return usedFields == maxUsedFields
}

func isWonSymbol(symbol string, table [3][3]string) bool {
	cnt := 0

	//смотрим горизонтально слева направо (построчно)
	for i := 0; i < 3; i++ {
		cnt = 0

		for j := 0; j < 3; j++ {
			if table[i][j] == symbol {
				cnt++
			}
		}

		if cnt == 3 {
			return true
		}
	}

	//смотрим вертикально сверху вниз (по колонкам)
	for j := 0; j < 3; j++ {
		cnt = 0

		for i := 0; i < 3; i++ {
			if table[i][j] == symbol {
				cnt++
			}
		}

		if cnt == 3 {
			return true
		}
	}

	//смотрим диагональ
	if table[0][0] == symbol && table[1][1] == symbol && table[2][2] == symbol || table[0][2] == symbol && table[1][1] == symbol && table[2][0] == symbol {
		return true
	}

	return false
}

func doTurn(table *[3][3]string, activePlayerNumber int, i int, j int) error {
	if table[i][j] != "~" {
		return errors.New("these indexes are already in use")
	}

	table[i][j] = PlayerSymbols[activePlayerNumber]

	return nil
}

func getIndexesFromPlayer() (int, int, error) {
	indexes := userInput("Enter i(row), j(column) indexes: ")
	indexes = strings.ReplaceAll(indexes, " ", "")

	i, err := strconv.Atoi(indexes[0:1])
	if err != nil {
		i = 0
	}

	j, err := strconv.Atoi(indexes[2:])
	if err != nil {
		j = 0
	}

	if i < 0 || i > 3 || j < 0 || j > 3 {
		return 0, 0, errors.New("indexes can be from 0 to 3")
	}

	return i, j, nil
}

func printTurnPlayer(playerName string) {
	fmt.Println(playerName + "'s turn.")
}

func fillTildaTable(table *[3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			table[i][j] = "~"
		}
	}
}

func drawTable(table [3][3]string) {
	fmt.Println("")
	fmt.Println("  0, 1, 2")
	for i := 0; i < 3; i++ {
		fmt.Print(i)
		fmt.Print(" ")
		fmt.Println(table[i])
	}
	println("")
}

func userInput(consoleMessage string) string {
	println("")
	println(consoleMessage)
	value, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	value = strings.TrimSuffix(value, "\n")
	return strings.TrimSuffix(value, "\r")
}
