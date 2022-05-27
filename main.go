package main

func main() {
	var h chan int
	go func() {
		h = make(chan int)
		<-h
	}()
	h <- 1
}
