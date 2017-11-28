// 16.3 Given two lists (A and B) of unique strings, write a program to determine
// if A is a subset of B.  That is, check if all the elements from A are
// contained in B.

package main

import (
  "fmt"
  "time"
)

func main() {
  sliceALen := getSliceLen('A')
  sliceAVal := make([]string, sliceALen)
  sliceAVal = getSliceVals(sliceALen, 'A')
  sliceBLen := getSliceLen('B')
  sliceBVal := make([]string, sliceBLen)
  sliceBVal = getSliceVals(sliceBLen, 'B')

  fmt.Println("Slice A\t: ", sliceAVal)
  fmt.Println("Slice B\t: ", sliceBVal)
  if len(sliceAVal) >= len(sliceBVal) {
    fmt.Println("Slice A is not a subset of Slice B.")
  } else {
    // Implementation 1: O(A * B)
    start1 := time.Now()
    result1Flag := checkSubset1(sliceAVal, sliceBVal)
    elapsed1 := time.Since(start1)
    if result1Flag {
      fmt.Println("Implementation 1: Slice A is a subset of Slice B.")
    } else {
      fmt.Println("Implementation 1: Slice A is not a subset of Slice B.")
    }

    // Implementation 2: O(A + B)
    hashTable := make(map[string]bool)
    start2    := time.Now()
    hashTable = loadHashTable(sliceBVal)
    elapsed2  := time.Since(start2)
    start3    := time.Now()
    result2Flag := checkSubset2(sliceAVal, hashTable)
    elapsed3  := time.Since(start3)
    if result2Flag {
      fmt.Println("Implementation 2: Slice A is a subset of Slice B.")
    } else {
      fmt.Println("Implementation 2: Slice A is not a subset of Slice B.")
    }

    fmt.Printf("Implementation 1\t: %s\n", elapsed1)
    fmt.Printf("Implementation 2(a)\t: %s\n", elapsed2)
    fmt.Printf("Implementation 2(b)\t: %s\n", elapsed3)
  }
}

func getSliceLen(sVal rune) int {
  var i int
  fmt.Println("Please enter the length of slice",string(sVal),":")
  fmt.Scan(&i)
  return i
}

func getSliceVals(sLen int, alpha rune) []string {
  sliceVal := make([]string, sLen)
  fmt.Println("Please enter the values for slice",string(alpha),", separated by spaces: ")
  for i := 0; i < sLen; i++ {
      fmt.Scan(&sliceVal[i])
    }
  return sliceVal
}

func checkSubset1(aVal []string, bVal []string) bool {
  found := 0
  for i := 0; i < len(aVal); i++ {
    for j := 0; j < len(bVal); j++ {
      if aVal[i] == bVal[j] {
        found++
        break
      }
      continue
    }
  }
  if found < len(aVal) {
    return false
  } else {
    return true
  }
}

func loadHashTable(sVal []string) map[string]bool {
  hashT := make(map[string]bool)
  for i:= 0; i < len(sVal); i++ {
    hashT[sVal[i]] = true
  }
  return hashT
}

func checkSubset2(sVal []string, hashTable map[string]bool) bool {
  found := 0
  for i:= 0; i < len(sVal); i++ {
    if _, ok := hashTable[sVal[i]]; ok {
      found++
    }
  }
  if found < len(sVal) {
    return false
  } else {
    return true
  }
}
