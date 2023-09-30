package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches   []string
	waitgroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
)

// 매개변수 루트로 주어진 경로부터 filename 이름을 갖는 파일이 존재하는지 탐색
func fileSearch(root string, filename string) {
	fmt.Println("Searching in", root)
	files, _ := os.ReadDir(root)

	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			// 여러 고루틴이 동시에 전역 변수 matches에 접근할 때 발생할 수 있는 레이스 컨디션 제한
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		// 디렉토리일 경우 해당 디렉토리도 탐색 및 고루틴 생성 (재귀)
		if file.IsDir() {
			waitgroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
	waitgroup.Done()
}

func main() {
	waitgroup.Add(1)
	go fileSearch("/Users/hongjinhyeok", "concurrent_file_search.go")
	waitgroup.Wait()
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
}
