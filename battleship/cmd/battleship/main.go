package main

import (
	"fmt"
	"os"

	"github.com/gdey/july_challenge/battleship"
)

func newGame() battleship.Game {
	game := battleship.NewGame()
	game.PlaceShip(battleship.Submarine, "F2", false)
	game.PlaceShip(battleship.Battleship, "M2", true)
	game.PlaceShip(battleship.Destroyer, "K2", true)
	game.PlaceShip(battleship.Cruiser, "D8", true)
	game.PlaceShip(battleship.AircraftCarrier, "E5", false)
	return game
}
func printBoard(g battleship.Game) {
	g.PrintBoard(os.Stdout)
	fmt.Println("\nA coordinate takes the form of 'Column''Row', i.e. A1, A16, P1, or P16")

}

func main() {
	input := make([]string, 5, 5)
	game := newGame()
	fmt.Println("Welcome to Battleship!\nThe current board looks like")
	printBoard(game)
	for !game.Ended() {
		fmt.Printf("\nCurrent Round: %d \t Score: %d\nPlease, enter 5 coordinates: ", game.Round(), game.Score())
		_, err := fmt.Scanln(&input[0], &input[1], &input[2], &input[3], &input[4])
		if err != nil {
			fmt.Println("You did not enter 5 good coordinates, please try again.", err)
			continue
		}
		hits, sank, err := game.Guess(input)
		if err != nil {
			fmt.Println("You did not enter 5 good coordinates, please try again.", err)
			continue
		}
		for k, _ := range sank {
			fmt.Printf("You sank my %s!\n", k)
		}
		if len(hits) > 0 {
			fmt.Printf("%d of your volleys hit!\n", len(hits))
		}
		printBoard(game)
	}
	if game.DidWin() {
		fmt.Println("Yay you won, you sank all my ships!")
	} else {
		shipsl := game.ShipsLeft()
		fmt.Printf("Sorry you did not manage to sink all my ships! I still had %d ships left.\n", len(shipsl))
		for _, s := range shipsl {
			fmt.Println("\t", s)
		}
	}
	game.ShowBoard()
	game.PrintBoard(os.Stdout)
	fmt.Printf("\nYou final score is %d\n Thank you for playing battleship!\n", game.Score())

}
