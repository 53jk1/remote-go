package main

import (
"bufio"
"fmt"
"os"
)

func main() {
fmt.Print("Input result: ")
reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')
if err != nil {
return err
}

fmt.Println(input)
}
