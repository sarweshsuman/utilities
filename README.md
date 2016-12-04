# utilities

This package is intended to contain all the handy functions related to strings, maps etc.

Right now it contains one function which helps in concating strings.

It can be called like below

import (
  "github.com/sarweshsuman/utilities"
  "fmt"
)

func main(){
    str := utilities.ConcatStrings("Hello"," World!")
    fmt.Println(str) 
}

Output :-
Hello World!
