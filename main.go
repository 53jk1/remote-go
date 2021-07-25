package main

import (
"fmt"
"strings"
)

func main() {
broken := "G# is g##d language!"
replacer := strings.NewReplacer("#", "o")
fixed := replacer.Replace(broken)
fmt.Println(fixed)
}
