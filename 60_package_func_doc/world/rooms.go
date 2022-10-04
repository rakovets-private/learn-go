package world

import "fmt"

const StartingRoom = "lobby"
const startingLevel = 1

func GetStartingLevel() int {
	return startingLevel
}

func chechStartingLevel(level int) bool {
	return level == startingLevel
}

func PrintStartRoom() {
	fmt.Println("Start level is", StartingRoom)
}
