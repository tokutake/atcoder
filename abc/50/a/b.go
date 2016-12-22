package main

import "fmt"

type Drink struct {
  ProblemIndex int
  Time int
}

func main() {
  var numberOfProblems int
  fmt.Scanf("%d", &numberOfProblems)
  solvingTimes  := make([]int, numberOfProblems)
  for i := 0; i < numberOfProblems; i++ {
    var time int
    fmt.Scanf("%d", &time)
    solvingTimes[i] = time
  }

  var numberOfDrinks int
  fmt.Scanf("%d", &numberOfDrinks)
  for i := 0; i < numberOfDrinks; i++ {
    var problemIndex, time int
    fmt.Scanf("%d %d", &problemIndex, &time)
    totalTime := 0
    for j := 0; j < numberOfProblems; j++ {
      if j == problemIndex - 1 {
        totalTime = totalTime + time
      } else {
        totalTime = totalTime + solvingTimes[j]
      }
    }
    fmt.Printf("%d\n", totalTime)
  }
}

