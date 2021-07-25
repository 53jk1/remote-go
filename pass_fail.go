package main

import (
"fmt"
"log"
)

func main() {
var i int
fmt.Print("Input result: ")
_, err := fmt.Scanf("%d", &i)
if err != nil {
log.Fatal(err)
}
if i == 100 {
fmt.Println("Great")
} else if i >= 60 {
fmt.Println("Passed")
} else {
fmt.Println("Failed")
}
}
