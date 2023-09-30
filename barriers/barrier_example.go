package main

import "time"

func waitOnBarrier(name string, timeToSleep int, barrier *Barrier) {
	for {
		println(name, "running")
		time.Sleep(time.Duration(timeToSleep) * time.Second)
		println(name, "is waiting on barrier")
		barrier.Wait()
	}
}

func main() {
	barrier := NewBarrier(2)
	go waitOnBarrier("red", 4, barrier)          // 4초마다 실행
	go waitOnBarrier("blue", 10, barrier)        // 10초마다 실행
	time.Sleep(time.Duration(100) * time.Second) // 100초 동안

}
