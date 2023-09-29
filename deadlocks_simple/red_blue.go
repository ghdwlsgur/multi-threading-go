package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

// 데드락은 두 개 이상의 프로세스나 스레드가 서로가 소유한 자원을 기다리며 영원히 차단되는 상황

// lock1, lock2로 임계영역 보호
// 하지만 lock1 먼저 잠금
func blueRobot() {
	for {
		fmt.Println("Blue: Acquiring lock1")
		lock1.Lock() // acquire lock1 first
		fmt.Println("Blue: Acquiring lock2")
		lock2.Lock() // acquire
		fmt.Println("Blue: Both locks Acquired")

		lock1.Unlock() // release
		lock2.Unlock() // release
		fmt.Println("Blue: Locks Released")
	}
}

// lock1, lock2로 임계영역 보호는 동일
// 하지만 lock2 먼저 잠금
func redRobot() {
	for {
		fmt.Println("Red: Acquiring lock2")
		lock2.Lock() // acquire lock2 first not lock1!
		fmt.Println("Red: Acquiring lock1")
		lock1.Lock() // acquire
		fmt.Println("Red: Both locks Acquired")

		lock1.Unlock() // release
		lock2.Unlock() // release
		fmt.Println("Red: Locks Released")
	}
}

func main() {
	// 두 고루틴이 동일한 순서로 뮤텍스를 획득하지 않기 때문에 데드락 발생 가능
	go redRobot()
	go blueRobot()

	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}
