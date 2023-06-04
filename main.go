package main

import (
	"fmt"
	"volleygame/services"
)

func main() {
	fmt.Println("\nVOLLEYBALL GAME")
	team := services.CreateTeam()
	game := team.StartGame()
	fmt.Println()
	fmt.Println("New game:", game.Team, "vs", game.Opponent)
	fmt.Println()
	game.Play()
}
