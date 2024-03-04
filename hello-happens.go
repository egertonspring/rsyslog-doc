package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Öffne die Log-Datei im Append-Modus
	f, err := os.OpenFile("logfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Setze den Logger auf die Log-Datei
	//logger := log.New(os.Stdout, "hello-go", log.LstdFlags)

	// Starte die Schleife, die alle 3 Sekunden läuft
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	randomTexts := []string{"Shit", "Book", "Chair", "Table", "Dog", "Cat", "House", "Car", "Tree",
		"Flower", "Computer", "Phone", "Lamp", "Door", "Window", "Pen", "Pencil", "Paper", "Desk",
		"Bed", "Cup", "Plate", "Spoon", "Fork", "Knife", "Shirt", "Pant", "Shoe", "Hat", "Jacket",
		"Coat", "Bag", "Wallet", "Watch", "Clock", "Key", "Guitar", "Piano", "Music", "Painting",
		"Brush", "Canvas", "Camera", "Film", "Photo", "Mirror", "Television", "Radio", "Microphone",
		"Speaker", "Headphone"}
	randomNames := []string{"Happens", "Runs", "Walks", "Jumps", "Eats", "Sleeps", "Reads", "Writes", "Talks", "Listens", "Sings",
		"Dances", "Swims", "Drives", "Flies", "Cooks", "Bakes", "Cleans", "Washes", "Brushes",
		"Combs", "Paints", "Draws", "Sketches", "Sculpts", "Builds", "Creates", "Plays", "Works",
		"Studies", "Learns", "Teaches", "Understands", "Remembers", "Forgets", "Believes", "Hopes",
		"Dreams", "Achieves", "Succeeds", "Fails", "Helps", "Supports", "Encourages", "Comforts", "Loves",
		"Hates", "Likes", "Dislikes", "Enjoys", "Relaxes"}

	for {
		select {
		case <-ticker.C:
			// Wähle zufällige Grußformel und Namen
			randomGreeting := randomTexts[rand.Intn(len(randomTexts))]
			randomName := randomNames[rand.Intn(len(randomNames))]

			// Erstelle den zufälligen Gruß
			randomGreetingText := fmt.Sprintf("%s %s!", randomGreeting, randomName)

			// Gib den Text auf stdout aus
			fmt.Println(randomGreetingText)

			// Logge den Text
			//logger.Println(randomGreetingText)
		}
	}
}
