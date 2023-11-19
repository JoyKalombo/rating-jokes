package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Joke represents a joke along with its total points
type Joke struct {
	Text   string
	Points int
}

// This function selects a random joke from the set of jokes I have
func getRandomJoke(jokes map[string][]Joke) Joke {

	// Select a random joke category
	categories := make([]string, 0, len(jokes))
	for category := range jokes {
		categories = append(categories, category)
	}
	selectedCategory := categories[rand.Intn(len(categories))]

	// Select a random joke from the chosen category
	jokeList := jokes[selectedCategory]
	selectedJoke := jokeList[rand.Intn(len(jokeList))]

	return selectedJoke
}

// I have broken down the code into smaller functions to improve readability and maintability.
func displaySummary(humourLevels int) {

	if humourLevels >= 0 && humourLevels <= 5 {
		fmt.Println("Your humour levels are very modest. Can we make you laugh more?")
	} else if humourLevels >= 6 && humourLevels <= 15 {
		fmt.Println("Your humour levels are moderate. You could laugh more, but you'd also be fine if you didn't.")
	} else if humourLevels >= 16 {
		fmt.Println("Your humour levels are abundant. You have reached the JOY level hahahaha")
	}

}

// This function will display the jokes and points from most popular to least popular
func displayJokesByPoints(jokePoints map[string]int) {
	// This will create a slice of type struct to hold the key-value pairs
	var sortedJokes []struct {
		Text   string
		Points int
	}

	// We are adding key-value pairs to the slice
	for text, points := range jokePoints {
		sortedJokes = append(sortedJokes, struct {
			Text   string
			Points int
		}{Text: text, Points: points})
	}
	// This will do the actual sorting in descending order of points
	sort.Slice(sortedJokes, func(i, j int) bool {
		return sortedJokes[i].Points > sortedJokes[j].Points
	})
	// This will display the sorted map
	fmt.Println("\nJokes and Points (Sorted, Most to least popular):")
	for _, joke := range sortedJokes {
		fmt.Printf("%s - Points: %d\n", joke.Text, joke.Points)
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	jokes := map[string][]Joke{
		"Programming": {
			{Text: "Why do programmers prefer dark mode? Because light attracts bugs! ", Points: 0},
			{Text: "\n What's a programmer's favorite snack? Code-nuts! ", Points: 0},
			{Text: "\n Why don't programmers like to go outside? The sunlight causes too many reflections! ", Points: 0},
		},
		"Puns": {
			{Text: "\n I used to be a baker because I kneaded dough! ", Points: 0},
			{Text: "\n I told my computer I needed a break, and now it won't stop sending me vacation ads! ", Points: 0},
			{Text: "\n I'm on a whiskey diet. I've lost three days already! ", Points: 0},
		},
		"Random": {
			{Text: "\n Why don't scientists trust atoms? Because they make up everything! ", Points: 0},
			{Text: "\n What did the ocean say to the shore? Nothing, it just waved! ", Points: 0},
			{Text: "\n Why did the scarecrow win an award? Because he was outstanding in his field! ", Points: 0},
		},
		"Sports/snacks": {
			{Text: "\n How do football players stay cool during the game? They stand near the fans! ", Points: 0},
			{Text: "\n What do you call cheese that isn't yours? Nacho cheese! ", Points: 0},
			{Text: "\n Why don't basketball players go on vacation? Because they would get called for travelling! ", Points: 0},
		},
	}

	jokePoints := make(map[string]int)
	humour_levels := 0

	fmt.Println("Welcome to the Silly Joke Generator!")

	for {
		fmt.Print("\nPress Enter to get a joke (or type 'exit' to quit): ")
		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			displaySummary(humour_levels)
			displayJokesByPoints(jokePoints)
			break
		}

		// This will get a random joke
		selectedJoke := getRandomJoke(jokes)
		fmt.Printf("\nCategory: %s\nJoke: %s\n", selectedJoke.Text, selectedJoke.Points)

		// This code allows users to rate the joke and incorporate error handling to ensure they only input valid numbers
		fmt.Print("\nOn a scale of 1 to 5, how funny was the joke?): ")
		var jokeLevel int
		_, err := fmt.Scanln(&jokeLevel)
		if err != nil {
			fmt.Println("Invalid input. Please enter a whole number from 1 to 5")
		}

		// This will update the total points for the joke and update the humour leves :-P
		fmt.Scanln(&jokeLevel)
		selectedJoke.Points += jokeLevel
		humour_levels += jokeLevel

		// This will update the map with the total points for the joke
		jokePoints[selectedJoke.Text] += jokeLevel
	}
}
