package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"`
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"`
	Comments   []Comment `json:"comments"`
}

func main() {
	doc := `
	[{
		"id": 1,
		"Title": "Hello, World!",
		"Author": {
			"Name": "Maria",
			"Email": "test@test.com"
		},
		"Content": "Hello~",
		"Recommends": ["John", "Andrew"],
		"Comments": [{
			"Id": 1,
			"Author": {
				"Name": "Andrew",
				"Email": "test2@hello.com"
			},
			"Content": "Hello Maria"
		}]
	}]
	`

	var data []Article

	err := json.Unmarshal([]byte(doc), &data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)

	newDoc, _ := json.MarshalIndent(data, "", " ")

	fmt.Println(string(newDoc))
}
