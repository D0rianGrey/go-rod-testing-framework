package main

import (
	"fmt"
	"time"
)

func sayHello(done chan bool) {
	for i := 0; i < 3; i++ {
		fmt.Println("Привет из горутины")
		time.Sleep(100 * time.Millisecond)
	}
	done <- true // сигнал о завершении работы
}

func main() {
	done := make(chan bool)
	go sayHello(done) // запуск горутины

	for i := 0; i < 3; i++ {
		fmt.Println("Привет из main")
		time.Sleep(150 * time.Millisecond)
	}
	<-done // ожидание сигнала завершения горутины
}
