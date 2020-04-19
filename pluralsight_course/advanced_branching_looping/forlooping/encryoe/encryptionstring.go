package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello To Encryption And Decryption")
	var s = func(value int) string {
		return "Anonymous Function" + strconv.Itoa(value)
	}(4)
	fmt.Println(s)

	// Channels For Message Passing
	stringChan := make(chan string) // Data
	towerChan1 := make(chan string) // Status Sender
	towerChan2 := make(chan string) // Status Receiver

	var offset int64 = 3

	go tower1(stringChan, towerChan1, offset)
	go tower2(stringChan, towerChan2, offset)

	for i := 0; i < 2; i++ {
		select {
		case towerObject := <-towerChan1:
			fmt.Printf("\nTower Object for Channel 1 is : %v\n", towerObject)
		case towerObject := <-towerChan2:
			fmt.Printf("\nTower Object for Channel 2 is : %v\n", towerObject)
		}
	}

}

func tower1(stringChan chan string, chan1 chan string, offset int64) {
	inputStream := bufio.NewReader(os.Stdin)
	fmt.Println("Tower 1 : Enter User Input Message for Tower 2 : ")
	userInput, _ := inputStream.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	fmt.Printf("\nTower 1 : Original String ==> %s", userInput)

	var secretString string
	for _, char := range userInput {
		secretString += string(int64(char) + offset)
	}

	fmt.Printf("\nTower 1 : Encrypted Secret String : %s", secretString)
	// Sending Data to Channel
	stringChan <- secretString
	chan1 <- "Tower 1 Message Sent With Encryption To Tower 2"
	//close(chan1) // Closing the channel
}

func tower2(stringChan chan string, chan2 chan string, offset int64) {

	// Receiving Data from Channel
	receivedMessage := <-stringChan
	var originalMessage string

	for _, char := range receivedMessage {
		originalMessage += string(int64(char) - offset)
	}
	fmt.Printf("\nTower 2 : Original Message From Tower 1 After Decryption is : %s", originalMessage)
	chan2 <- "Tower 2 : Message Received From Tower 1 Successfully"
}
