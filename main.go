package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	fmt.Println(CreateHelloString())
}

func CreateHelloString() (message string) {
	message = emoji.Sprint("Hello :world_map:!")
	return
}
