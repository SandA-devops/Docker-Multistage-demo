package main

import (
  "fmt"
  "os"
)

func main() {
	fmt.Println("Example go app - I dont do much")
	fmt.Println(os.Getenv("MESSAGE_VAR"))
}
