// 16.2 Reverse the order of elements in an array (without creating a new array)

package main

import "fmt"

func main() {
  var sliceLen int
  sliceVal := make([]int, sliceLen)

  sliceLen = getSliceLen()
  sliceVal = getSliceVals(sliceLen)

  fmt.Println("Existing slice\t: ", sliceVal)
  reverseSlice(sliceVal)
  fmt.Println("Reversed slice\t: ", sliceVal)
}

func getSliceLen() int {
  var i int
  fmt.Println("Please enter the length of the slice: ")
  fmt.Scan(&i)
  return i
}

func getSliceVals(sVal int) []int {
  sliceVal := make([]int, sVal)
  fmt.Println("Please enter the ordered slice values separated by spaces: ")
  for i := 0; i < sVal; i++ {
      fmt.Scan(&sliceVal[i])
    }
  return sliceVal
}

func reverseSlice(sVal []int) {
  tempVal := 0
  for i := 0; i < len(sVal) / 2; i++ {
    tempVal                 = sVal[i]
    sVal[i]                 = sVal[len(sVal) - 1 - i]
    sVal[len(sVal) - 1 - i] = tempVal
  }
}
