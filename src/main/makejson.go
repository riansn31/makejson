package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	var address string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your address: ")
	fmt.Scanln(&address)
	data := map[string]string{
		"name":    name,
		"address": address,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	dataBytes, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println(string(dataBytes))
}
