package main
import (
	"fmt"
	"log"
	"mkgo/internal/engine"
)

func main() {
	fmt.Println("Starting MKGO")
	game:= engine.New()
	if err := game.Run(); err != nil {
		log.Fatal("Error running game:", err)
	}
	fmt.Println("Goodbye")
}
