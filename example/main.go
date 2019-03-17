package main

import (
	"fmt"
	"time"

	"github.com/toshusai/bui/view"
)

func main() {
	w := view.NewWindow(800, 600, "Test")
	t := time.Now()
	w.Update = func() {
		timer := time.Now().Sub(t)
		fmt.Println(timer)
		if timer.Seconds() > 3 {
			w.Close()
		}
	}
	w.Run()
}
