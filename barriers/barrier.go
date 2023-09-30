package main

import "sync"

type Barrier struct {
	total int         // 고루틴의 총 개수
	count int         // Barrier를 기다리고 있는 고루틴의 개수
	mutex *sync.Mutex // 스레드 동기화를 위한 뮤텍스
	cond  *sync.Cond  // 조건 변수
}

// Barrier 생성 및 초기화
func NewBarrier(size int) *Barrier {
	lockToUse := &sync.Mutex{}
	condToUse := sync.NewCond(lockToUse)
	return &Barrier{size, size, lockToUse, condToUse}
}

// 모든 고루틴이 들어올 때까지 대기
func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count -= 1

	if b.count == 0 {
		b.count = b.total
		// Broadcast wakes all goroutines waiting on c.
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}

	b.mutex.Unlock()
}
