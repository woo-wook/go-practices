package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	data := make([]Article, 1)

	data[0].Id = 1
	data[0].Title = "Hello, World!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@text.com"
	data[0].Content = "Hello!"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew.@text.com"
	data[0].Comments[0].Content = "Hello, Maria!"

	doc, _ := json.Marshal(data)

	err := os.WriteFile("src/json/articles.json", doc, os.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		return
	}

	readDoc, err := os.ReadFile("src/json/articles.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	var readData []Article

	json.Unmarshal(readDoc, &readData)

	fmt.Println(readData)
}
