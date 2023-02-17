package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	title := flag.String("title", "", "영화 명")
	runtime := flag.Int("runtime", 0, "상영 시간")
	rating := flag.Float64("rating", 0.0, "평점")

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	fmt.Printf("영화 이름: %s\n상영 시간: %d분\n평점: %f", *title, *runtime, *rating)
	fmt.Println(os.Args)
}
