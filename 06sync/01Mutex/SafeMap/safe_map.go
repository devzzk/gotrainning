package SafeMap

import "sync"

type SafeMap struct {
	mutex sync.RWMutex
	// Mutex he RWMutex 都是不可重入的，
	// 同一个 goroutin 容易死锁，
}

func (s *SafeMap) CheckAndDoSomething() {
	s.mutex.Lock()
	// check and do something
	// 适合写多读少的场景

	s.mutex.Unlock()
}

func (s *SafeMap) CheckAndDoSomething1() {
	s.mutex.RLock()
	// check 第一次检查
	// 读多写少的场景
	s.mutex.RUnlock()

	s.mutex.Lock()
	// check and do something
	defer s.mutex.Unlock()
	// panic 是直接中断， 所以如果此时拿着锁， 那么是无法释放掉的，
	// 但defer可以确保即使panic 了， 这个动作也会执行
}
