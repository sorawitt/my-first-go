package sunwukong

import (
	"fmt"

	"github.com/sorawitt/my-first-go/sunwukong/internal/kong"
)

func SayHelloSunWuKong() {
	fmt.Println("Hello SunWuKong")
	kong.SayHelloKong()
}
