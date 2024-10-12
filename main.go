package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sorawitt/my-first-go/sunwukong"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello world")
	fmt.Printf("UUID: %s", id)
	sunwukong.SayHelloSunWuKong()
}
