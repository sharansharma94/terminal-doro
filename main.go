package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	fmt.Printf("Welcome to terminal doro ")

	var option string
	for {

		fmt.Printf("Choose Options \n")
		fmt.Printf("s : start \n")
		fmt.Printf("e : stop \n")
		fmt.Printf("p : pause \n")

		fmt.Scan(&option)

		switch option {
		case "s":
			_, err := startTimer()
			if err != nil {
				log.Fatalf(err.Error())
			}

			// fmt.Println("retimer", option)
		case "e":
			stopTimer()
		case "p":
			pauseTimer()
		case "r":
			resumeTimer()
		}

		fmt.Println("you choosed:", option)
	}
}

func startTimer() (time.Duration, error) {
	fmt.Println("Timer started")

	date := "1h2m3s"

	targetTime, err := time.Parse(time.Kitchen, date)

	if err != nil {
		return time.Duration(0), err
	}

	fmt.Println("target time :", targetTime)
	return time.Duration(1), nil
}
func stopTimer() {
	fmt.Println("Timer stoped")
}
func pauseTimer() {
	fmt.Println("Timer paused")
}
func resumeTimer() {
	fmt.Println("Timer resumed")
}
