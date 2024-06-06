package main

import (
	"fmt"
	"simple-spinner/animation"
	"simple-spinner/spinner"
	"time"
)

func main() {
	s := spinner.NewSpinner("Processing ", animation.Dots)

	s.Start()
	defer s.Stop()
	time.Sleep(3 * time.Second)
	fmt.Print("Completed\n")
}
