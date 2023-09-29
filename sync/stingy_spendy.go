package main

import (
	"sync"
	"time"
)

var (
	money = 100
	lock  = sync.Mutex{}
)

func stingy() {
	println("Stingy Start")

	// 1부터 1000까지 전역 변수 money를 10씩 증가
	for i := 1; i <= 1000; i++ {
		// 상호 배제를 위한 뮤텍스 잠금
		lock.Lock()
		money += 10
		// 뮤텍스 잠금 해제
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy Done")
}

func spendy() {
	println("Spendy Start")

	// 1부터 1000까지 전역 변수 money를 10씩 감소
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money -= 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Spendy Done")
}

func main() {
	// 고루틴으로 두 함수 동시 실행
	// 뮤텍스 잠금으로 전역 변수에 두 함수가 상호배제하도록 하였지만 실행 순서는 제어하지 않았다.
	go stingy()
	go spendy()

	time.Sleep(3000 * time.Millisecond)
	print(money)
}
