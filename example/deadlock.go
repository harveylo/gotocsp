package main

import (
	"fmt"
	"sync"

	csp "github.com/harveylo/gotocsp"
)

func main() {
	wg := sync.WaitGroup{}
	ch0 := csp.MakeChannel(0)
	ch1 := csp.MakeChannel(0)
	args := []int{1, 2, 3}
	csp.MakeProcessWithChannel(f1, args, &wg, ch0)
	csp.MakeProcessWithChannel(f2, args, &wg, ch1)
	csp.ChannelOutput(ch1, args[1])
	v := csp.ChannelInput(ch0)
	csp.EndMainProcess(&wg)
	fmt.Println(v)
}

func f1(args []int, wg *sync.WaitGroup, ch chan int) {
	v := csp.ChannelInput(ch)
	csp.EndProcessNotMain(wg)
	fmt.Println(v)
}

func f2(args []int, wg *sync.WaitGroup, ch chan int) {
	csp.ChannelOutput(ch, args[0])
	csp.EndProcessNotMain(wg)
}
