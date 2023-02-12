package main

import "fmt"

type Duck struct {
}

type Person struct {
}

type Quacker interface {
	quack()
	feathers()
}

func main() {
	var donald Duck
	var john Person

	inTheForest(donald)
	inTheForest(john)

	_, ok := interface{}(john).(Quacker)
	fmt.Println(ok)
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}

func (d Duck) quack() {
	fmt.Println("꽤에에에엑!")
}

func (d Duck) feathers() {
	fmt.Println("오리는 흰색과 회색털을 가지고 있다.")
}

func (p Person) quack() {
	fmt.Println("사람을 오리를 흉내냅니다. 꽤에에에엑!")
}

func (p Person) feathers() {
	fmt.Println("깃털을 땅에서 주워서 보여준다.")
}
