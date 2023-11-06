package main

import (
	"fmt"
	"sync"

	csp "github.com/harveylo/gotocsp"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	wg := sync.WaitGroup{}
	m := 3
	length := len(s)
	ch := csp.MakeChannel(0)
	for i := 0; i < m; i++ {
		csp.MakeProcessWithChannel(calSum, s[length/m*i:length/m*(i+1)], &wg, ch)
	}
	sum := 0
	for i := 0; i < m; i++ {
		v := csp.ChannelInput(ch)
		sum += v
	}
	csp.EndMainProcess(&wg)
	fmt.Println(sum)
}

func calSum(s []int, wg *sync.WaitGroup, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	csp.ChannelOutput(ch, sum)
	csp.EndProcessNotMain(wg)
}
