package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	fmt.Print("enter a name: ")
	fmt.Scan(&name)
	fmt.Print("enter an address: ")
	fmt.Scan(&address)
	inputs := map[string]string{"name": name, "address": address}
	json, _ := json.Marshal(inputs)
	fmt.Println(string(json))
}
