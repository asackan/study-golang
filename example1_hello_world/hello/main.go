package main

import (
	"fmt"

	"github.com/asackan/study-golang/example1_hello_world/greeting"
)

func main() {
	message := greeting.Hello("John")
	fmt.Println(message)
}
