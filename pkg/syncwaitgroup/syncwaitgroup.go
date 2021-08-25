package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	i int64
	mu sync.Mutex // 공유 데이터 i를 보호하기 위한 뮤텍스
}

// counter 값을 1씩 증가시킴
func (c *counter) increment() {
	c.mu.Lock()   // i 값을 변경하는 부분(임계 영역)을 뮤텍스로 잠금
	c.i += 1      // 공유 데이터 변경
	c.mu.Unlock() // i 값을 변경 완료한 후 뮤텍스 잠금 해제
}
func (c *counter) decrement() {
	c.mu.Lock()   // i 값을 변경하는 부분(임계 영역)을 뮤텍스로 잠금
	c.i -= 1      // 공유 데이터 변경
	c.mu.Unlock() // i 값을 변경 완료한 후 뮤텍스 잠금 해제
}


// counter의 값을 출력
func (c *counter) display() {
	fmt.Println(c.i)
}

func main2() {
	// 모든 CPU를 사용하게 함
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}     // 카운터 생성
	wg := sync.WaitGroup{} // WaitGroup 생성

	// c.increment()를 실행하는 고루틴 1000개 실행

	wg.Add(5) // WaitGroup의 고루틴 개수 1 증가
	go func() {
		defer wg.Done() // 고루틴 종료 시 Done() 처리
		go func() {
			defer wg.Done() // 고루틴 종료 시 Done() 처리
			c.increment()
		}()
		c.increment()   // 카운터 값을 1 증가시킴
	}()
	go func() {
		defer wg.Done() // 고루틴 종료 시 Done() 처리
		go func() {
			defer wg.Done() // 고루틴 종료 시 Done() 처리
			c.decrement()
			go func() {
				defer wg.Done() // 고루틴 종료 시 Done() 처리
				c.decrement()
			}()
		}()
		c.decrement()   // 카운터 값을 1 증가시킴
	}()

	wg.Wait()   // 모든 고루틴이 종료될 때까지 대기

	c.display() // c의 값 출력
}
