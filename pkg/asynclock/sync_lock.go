package asynclock

import "sync"

type SyncLock struct {
	mu     sync.Mutex
	counts map[string]int
	locks  map[string]*sync.Mutex
}

func NewSycnLock() *SyncLock {
	return &SyncLock{
		counts: make(map[string]int),
		locks:  make(map[string]*sync.Mutex),
	}
}

func (syncLock *SyncLock) Lock(key string) {
	syncLock.mu.Lock()
	sylock, exists := syncLock.locks[key]
	if !exists {
		sylock = &sync.Mutex{}
		syncLock.locks[key] = sylock
	}

	syncLock.counts[key]++
	syncLock.mu.Unlock()

	sylock.Lock()
}

func (syncLock *SyncLock) Unlock(key string) {
	syncLock.mu.Lock()
	defer syncLock.mu.Unlock()

	if syLock, exists := syncLock.locks[key]; exists {
		syLock.Unlock()
		syncLock.counts[key]--

		if syncLock.counts[key] == 0 {
			delete(syncLock.locks, key)
			delete(syncLock.counts, key)
		}
	}
}
