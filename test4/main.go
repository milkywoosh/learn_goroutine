package m

import (
	"fmt"
	"runtime"
)

// this func receive chan string
func Printer(data chan string) {
	x := <-data
	fmt.Printf("%v\n", x)

}

var slice_of_data []string = []string{"one", "two", "three"}
var chan_receiver = make(chan string)

func Looper(person string) {
	s := fmt.Sprintf("hello %s", person)
	chan_receiver <- s
}

func main() {
	runtime.GOMAXPROCS(1)

	// slice_of_data := []string{"one", "two", "three"}
	for _, data := range slice_of_data {
		go Looper(data)
	}

	for i := 0; i < 3; i++ {
		Printer(chan_receiver)
	}

}
