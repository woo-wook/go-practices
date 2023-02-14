package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	doc := `
	{
		"name": "maria",
		"age": 10
	}
	`

	var data map[string]interface{}

	json.Unmarshal([]byte(doc), &data)

	fmt.Println(data["name"], data["age"])

	newDoc, _ := json.Marshal(data)
	fmt.Println(string(newDoc))

	newDoc2, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(newDoc2))
}
