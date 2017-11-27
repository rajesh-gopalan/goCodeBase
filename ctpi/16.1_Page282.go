package main

import "fmt"

func main() {
  var sliceLen int
  sliceVal := make([]int, sliceLen)
  var newElement int

  sliceLen = getSliceLen()
  sliceVal = getSliceVals(sliceLen)
  newElement = getNewElem()

  fmt.Println("Existing slice value\t: ", sliceVal)
  var successFlag bool = insertNewElem(sliceVal, newElement)
  if successFlag {
    fmt.Println("New slice value\t\t: ", sliceVal)
  } else {
    fmt.Println("Erorr: Please satisfy requirements")
  }
}

// Supporting functions

func getSliceLen() int {
  var i int
  fmt.Println("Enter the length of the slice: ")
  fmt.Scan(&i)
  return i
}

func getSliceVals(sVal int) []int {
  sliceVal := make([]int, sVal)
  fmt.Println("Please enter the positive, sorted slice values separated by spaces (last element must be a zero): ")
  for i := 0; i < sVal; i++ {
      fmt.Scan(&sliceVal[i])
    }
  return sliceVal
}

func getNewElem() int {
  var nElem int
  fmt.Println("Please enter the new element to be inserted: ")
  fmt.Scan(&nElem)
  return nElem
}

func insertNewElem(sVal []int, newElem int) bool {
  var loopVar int = len(sVal) - 1
  if sVal[loopVar] != 0 || newElem <= 0  {
    return false
  }
  for loopVar >= 0 && sVal[loopVar - 1] > newElem {
    sVal[loopVar]     = sVal[loopVar - 1]
    sVal[loopVar - 1] = newElem
    loopVar--
  }
  return true
}
