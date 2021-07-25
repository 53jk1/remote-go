package main

import (
"bufio"
"fmt"
"os"
"log"
)

func main() {
fmt.Print("Input result: ")
reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')
if err != nil {
log.Fatal(err)
}

fmt.Println(input)
}
