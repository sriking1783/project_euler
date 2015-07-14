package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "math/big"
)



func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
  dat, err := ioutil.ReadFile("/tmp/numbers")
  check(err)
  contents := strings.Split(string(dat),"\n")
  sum := big.NewInt(0)
  fmt.Println(sum)
  bi := big.NewInt(0)
  for i := 0; i < 100; i++ {
    if _, ok := bi.SetString(contents[i], 10); ok {
      //fmt.Printf("%v\n", bi)
      sum.Add(sum, bi)
    } else {
      fmt.Printf("error parsing line %#v\n", i)
    }
  }
  fmt.Println(sum.String()[0:10])
}
