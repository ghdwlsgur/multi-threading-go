package main

import (
	"sync/atomic"
	"time"
)

var (
	money int32 = 100
)

// 고루틴에서 접근 가능한 원자성 카운터를 사용
func stingy() {
	for i := 1; i <= 1000; i++ {
		// 원자적으로 값 증가, 매개변수로 증감값과 메모리 주소 전달
		atomic.AddInt32(&money, 10)
		time.Sleep(1 * time.Millisecond)
		println("stindy", money)
	}
	println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, -10)
		time.Sleep(1 * time.Millisecond)
		println("spendy", money)
	}
	println("Spendy Done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}
