package main

import "fmt"

func main() {
	animalChannel := make(chan string)
	fruitChannel := make(chan string)

	go func() {
		defer close(animalChannel)
		animalChannel <- "Duck"
		animalChannel <- "Chicken"
		animalChannel <- "Bird"
		animalChannel <- "Monkey"
		animalChannel <- "Cow"
	}()

	go func() {
		defer close(fruitChannel)
		fruitChannel <- "Apple"
		fruitChannel <- "Banana"
		fruitChannel <- "Orange"
		fruitChannel <- "Tomato"
		fruitChannel <- "Grape"
	}()

	fmt.Println("Proccess")

	var animalStatus bool
	var fruitStatus bool

	for {
		select {
		case data, status := <-animalChannel:
			animalStatus = status
			if animalStatus {
				fmt.Println("Animal : " + data)
			}
		case data, status := <-fruitChannel:
			fruitStatus = status
			if fruitStatus {
				fmt.Println("Fruit : " + data)
			}
		}
		if (!animalStatus && !fruitStatus) {
			break
		}
	}


	fmt.Println("Done")
}