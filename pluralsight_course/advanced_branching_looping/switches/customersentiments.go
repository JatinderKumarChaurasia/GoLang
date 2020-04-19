package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type rating float64

// constants sentiments
const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	negative      rating = -0.6
	extraNegative rating = -1.2
	initial       rating = 5.0
)

type result struct {
	feedbackDate     string
	feedbackTotal    int
	feedbackPositive int
	feedbackNegative int
	feedbackNeutral  int
}

var customer []rating

func main() {
	inputFile, err := os.Open("data.csv")
	if err != nil {
		//log.Fatal("Unable to Open File ................ ")
		exitProgram("Unable to Open File ................ ", err.Error())
	}
	defer func() {
		corr := inputFile.Close()
		if err == nil {
			err = corr
		}
	}()
	read := bufio.NewReader(inputFile)
	str, err1 := read.ReadString('\n')
	if err1 != nil {
		exitProgram("Unable to read file", err1.Error())
	}
	var feedback result
	feedback.feedbackDate = str
	for {
		str, err := read.ReadString('\n')
		if err != nil {
			break
		}
		if len(str) > 10 {
			feedback.feedbackTotal++
			var customerRating rating
			customerRating = initial
			text := strings.Split(str, " ")
			for _, word := range text {
				switch s := strings.Trim(strings.ToLower(word), ",.,!,?,\t,\n,\r"); s {
				case "pleasure", "impressed", "wonderful", "fantastic", "splendid":
					customerRating += extraPositive
				case "help", "thanks", "helpful", "happy":
					customerRating += positive
				case "not helpful", "sad", "angry", "improve", "annoy":
					customerRating += negative
				case "pathetic", "bad", "worse", "agitated", "frustrated", "unfortunately":
					customerRating += extraNegative
				}
			}
			switch {
			case customerRating > 8.0:
				feedback.feedbackPositive++
			case customerRating >= 4.0 && customerRating <= 8.0:
				feedback.feedbackNeutral++
			case customerRating < 4.0:
				feedback.feedbackNegative++
			}
			customer = append(customer, customerRating)
		}
	}
	feedbackTable(feedback, customer)
}

func feedbackTable(feedback result, ratings []rating) {
	fmt.Printf("Report for Date: %s\n", feedback.feedbackDate)
	fmt.Printf("Total Customer Reviews : %d\n", feedback.feedbackTotal)
	fmt.Printf("Positive Reviews : %d\n", feedback.feedbackPositive)
	fmt.Printf("Negative Reviews : %d\n", feedback.feedbackNegative)
	fmt.Printf("Neutral Reviews : %d\n", feedback.feedbackNeutral)
	fmt.Println("Customer Ratings : ", ratings)
}

func exitProgram(msg string, exitError string) {
	log.Fatal(msg+"exiting the program with error :", exitError)
}
