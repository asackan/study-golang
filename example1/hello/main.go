package main

import (
	"fmt"

	"github.com/asackan/study-golang/example1/greeting"
)

func main() {
	message := greeting.Hello("John")
	fmt.Println(message)
}
