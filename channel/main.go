package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string
	Age  int
}

func printUser(u user) {
	myPf("name: %s, age: %v", u.Name, u.Age)
}

func myPf(fomart string, v ...any) {
	fmt.Printf("%s %s", time.Now().Format(time.DateTime), fmt.Sprintf(fomart, v...))
	fmt.Println()
}

func main() {

	userChan := make(chan user, 100)

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(2 * time.Second)
			myPf("i: %v", i)
			userChan <- user{
				Name: fmt.Sprintf("user%v", i),
				Age:  i,
			}
		}
	}()

	for {
		select {
		case u := <-userChan:
			time.Sleep(5 * time.Second)
			printUser(u)
		}
	}
}
