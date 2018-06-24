package main

import (
	"fmt"
	"time"
	"strconv"
)

var PizzaNum=0
var PizzaName=""


func makeDough(String chan string){
	PizzaNum++
	PizzaName="Pizza #"strconv.Itoa(pizzaNum)
}

func main() {
	for i := 0; i < 10; i++ {
		go count(i)
	}
	time.Sleep(time.Millisecond * 11000)
}
