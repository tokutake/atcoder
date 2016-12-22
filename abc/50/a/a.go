package main

import "fmt"

func main() {
  var operand1, operand2 int
  var operator string
  for {
    n, err := fmt.Scanf("%d %s %d", &operand1, &operator, &operand2)
    if n == 0 && err != nil {
      break
    }
    var answer int
    if operator == "+" {
      answer = operand1 + operand2
    } else {
      answer = operand1 - operand2
    }
    fmt.Printf("%d\n", answer)
  }
}

