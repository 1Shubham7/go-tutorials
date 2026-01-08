package main

import (
	"fmt"
	"time"
)

func main(){

	done := make(chan struct{})
	go func() {
		fmt.Println("we just entered the go routine")
		time.Sleep(1*time.Second)
		fmt.Println("task completed")
		done <- struct{}{}
	}()

	<-done

	fmt.Println("go routine complete")

}
