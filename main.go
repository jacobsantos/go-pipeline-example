package main

import "time"

func simulate(duration time.Duration, input, output, quit chan int) {
	for {
		select {
		case <-input:
			time.Sleep(duration)
			output <- 1
		case <-quit:
			quit <- 1
			return
		}
	}
}

func terminateAfter(after int, duration time.Duration, input, quit chan int) {
	at := 0
	for {
		select {
		case <-input:
			if at >= after {
				quit <- 1
				return
			}
			time.Sleep(duration)
			at++
		case <-quit:
			quit <- 1
			return
		}
	}
}

func start(duration time.Duration, output chan int) {
	for {
		output <- 1
		time.Sleep(duration)
	}
}

func main() {
	firstTask := make(chan int)
	secondTask := make(chan int)
	thirdTask := make(chan int)
	quit := make(chan int)
	go simulate(time.Second*2, firstTask, secondTask, quit)
	go simulate(time.Second, secondTask, thirdTask, quit)
	go terminateAfter(100, time.Second*3, thirdTask, quit)
	go start(time.Millisecond*50, firstTask)
	<-quit
}
