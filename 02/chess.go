package main

import (
  "fmt"
  "bytes"
)

func main() {
  str := ""
  sharp := "#" 
  black := " "
  // naive concate
  for i := 1; i <= 8; i++ {
    str = "" 
    for j := 1; j <= 8; j++ {
      if j % 2 == 0 {
        str += sharp
      } else {
        str += black
      }
    }
    sharp, black = black, sharp
    fmt.Println(str)
  }
  // append concate
  for i := 1; i <= 8; i++ {
    s := []string{}
    for j := 1; j <= 8; j++ {
      if j % 2 == 0 {
        s = append(s, sharp)
      } else {
        s = append(s, black)
      }
    }
    sharp, black = black, sharp
    fmt.Println(s)
  }
  // buffer
  for i := 1; i <= 8; i++ {
    var buffer bytes.Buffer
    for j := 1; j <= 8; j++ {
      if j % 2 == 0 {
        buffer.WriteString(sharp)
      } else {
        buffer.WriteString(black)
      }
    }
    sharp, black = black, sharp
    fmt.Println(buffer.String())
  }
}