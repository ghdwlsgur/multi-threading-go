package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	money          = 100                 // 공유 자원 (전역 변수)
	lock           = sync.Mutex{}        // 뮤텍스
	moneyDeposited = sync.NewCond(&lock) // 조건변수
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		fmt.Println("Stingy sees balance of ", money)
		// 대기 중인 고루틴에 공유 자원인 money 변수 값 변경 알림
		// Signal wakes one goroutine waiting on c
		moneyDeposited.Signal()
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		// money가 20 미만이면 대기
		for money-20 < 0 {
			moneyDeposited.Wait()
		}
		money -= 20
		fmt.Println("Spendy sees balance of ", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Spendy Done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}
