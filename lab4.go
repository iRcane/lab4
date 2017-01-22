package main

import (
	"fmt"
	"time"
)

type Token struct {
	data string
	recipient int
}

func initial(ch chan Token, token Token) {
	fmt.Println("initial thread")
	ch <- token
}

func thread(ch chan Token, num int) {
	fmt.Println("thread ", num) 
	token := <-ch
	if token.recipient == num {
		fmt.Println(token.data, "received")
	} else {
		ch <- token
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	token := Token {"some data", n}
	ch := make(chan Token)
	go initial(ch, token)
	for i := 1; i <= n; i++ {
		go thread(ch, i)
	}
	time.Sleep(time.Second)
}