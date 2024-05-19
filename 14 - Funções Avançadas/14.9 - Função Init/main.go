package main

import (
	"fmt"
	"time"
)

var weekday string

func init() {
	weekday = time.Now().Weekday().String()
}

func main() {

	fmt.Printf("Olá hoje é: %s\n", weekday)
}
