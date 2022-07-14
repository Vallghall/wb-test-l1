package main

import "time"

func main() {
	sleep(time.Second * 3)
}

// sleep func implementation
func sleep(n time.Duration) {
	<-time.After(n)
}
