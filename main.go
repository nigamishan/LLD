package main

import (
	"SnakesAndLadder/Manager"
	"SnakesAndLadder/entities"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func main() {
	// Take Config
	config := entities.Config{}
	file, _ := os.Open("config/config.yaml")
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatal("Unable to Decode")
	}

	// create Board
	board := entities.CreateBoard(config.BoardSize*config.BoardSize, config.DieCount, config.MoveStrategy)

	// Add snakes and ladders
	var s, l, start, end int
	s = config.SnakeCount
	l = config.LadderCount
	fmt.Printf("Enter %d Snakes\n", s)
	for i := 0; i < s; i++ {
		fmt.Scanf("%d %d", &start, &end)
		if !board.AddActors(start, end, "Snake") {
			fmt.Println("Incorrect entry, pls add again")
			i--
			continue
		}
	}
	fmt.Printf("Enter %d Ladders\n", l)
	for i := 0; i < l; i++ {
		fmt.Scanf("%d %d", &start, &end)
		if !board.AddActors(start, end, "Ladder") {
			fmt.Println("Incorrect entry, pls add again")
			i--
			continue
		}
	}

	// Create Game
	game := Manager.CreateGame(board)

	// create User
	var (
		userCount int
		name      string
	)
	userCount = config.PlayerCount
	fmt.Printf("Enter %d Player name\n", userCount)
	for i := 0; i < userCount; i++ {
		fmt.Scanf("%s", &name)
		game.AddPlayers(entities.CreateUser(1, name))
	}

	game.StartGame()

}
