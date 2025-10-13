package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	month := now.Month()
	var day int = now.Day()
	fmt.Println(month, day)

	univ := "go$ inha!"
	changer := strings.NewReplacer("$", "!")
	changed := changer.Replace(univ)
	fmt.Println(changed)
}
