package spinLock

import (
    "runtime"
    "sync/atomic"
)

// 自旋锁
type SpinLock struct {
    lock uintptr
}

// 加锁
func (l *SpinLock) Lock() {
    for !l.TryLock() {
        runtime.Gosched() // //allow other goroutines to do stuff.
        continue
    }
}

// 尝试加锁
func (l *SpinLock) TryLock() bool {
    return atomic.CompareAndSwapUintptr(&l.lock, 0, 1)
}

// 解锁
func (l *SpinLock) Unlock() {
    atomic.StoreUintptr(&l.lock, 0)
}
