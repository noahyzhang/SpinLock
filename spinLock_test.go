package spinLock

import (
    "sync"
    "testing"
)

// ----------- 当临界区仅仅为 i++ 时测试无锁、互斥锁、自旋锁的性能 ------------------

// 不加锁时且串行时，i++ 的性能
func BenchmarkAdd(b *testing.B) {
    number := 0
    for i := 0; i < b.N; i++ {
        number++
    }
}

// 不加锁且并行时，i++ 的性能
func BenchmarkParallelAdd(b *testing.B) {
    b.SetParallelism(50)
    number := 0
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            number++
        }
    })
}

// 测试串行情况下，临界区只为 i++ 时，自旋锁性能
func BenchmarkSpinLockAdd(t *testing.B) {
    loc := SpinLock{}
    number := 0
    for i := 0; i < t.N; i++ {
        loc.Lock()
        number++
        loc.Unlock()
    }
    if number != t.N {
        t.Errorf("Expected %d, but got %d\n", t.N, number)
    }
}

// 测试并行情况下，临界区只为 i++ 时，自旋锁性能
func BenchmarkParallelSpinLockAdd(b *testing.B) {
    loc := SpinLock{}
    b.SetParallelism(50)
    number := 0
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            loc.Lock()
            number++
            loc.Unlock()
        }
    })
    if number != b.N {
        b.Errorf("Expected %d, but got %d\n", b.N, number)
    }
}

// 测试串行情况下，临界区只为 i++ 时，互斥锁性能
func BenchmarkMutexAdd(t *testing.B) {
    loc := sync.Mutex{}
    number := 0
    for i := 0; i < t.N; i++ {
        loc.Lock()
        number++
        loc.Unlock()
    }
    if number != t.N {
        t.Errorf("Excepted %d, but got %d\n", t.N, number)
    }
}

// 测试并行情况下，临界区只为 i++ 时，互斥锁性能
func BenchmarkParallelMutexAdd(b *testing.B) {
    loc := sync.Mutex{}
    number := 0
    b.SetParallelism(50)
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            loc.Lock()
            number++
            loc.Unlock()
        }
    })
    if number != b.N {
        b.Errorf("Excepted %d, but got %d\n", b.N, number)
    }
}

// ----------- 当临界区为求斐波那契数列时测试无锁、互斥锁、自旋锁的性能 ------------------

// 斐波那契数列计算
func fib(n int) int {
    if n < 2 {
        return n
    }
    return fib(n-1) + fib(n-2)
}

// 测试串行情况，临界区为求斐波那契数列时，不加锁性能
func BenchmarkFib(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fib(10)
    }
}

// 测试并行情况，临界区为求斐波那契数列时，不加锁性能
func BenchmarkParallelFib(b *testing.B) {
    b.SetParallelism(50)
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            fib(10)
        }
    })
}

// 测试串行情况下，临界区为求斐波那契数列时，自旋锁性能
func BenchmarkSpinLockFib(b *testing.B) {
    loc := SpinLock{}
    for i := 0; i < b.N; i++ {
        loc.Lock()
        fib(10)
        loc.Unlock()
    }
}

// 测试并行情况下，临界区为求斐波那契数列时，自旋锁性能
func BenchmarkParallelSpinLockFib(b *testing.B) {
    loc := SpinLock{}
    b.SetParallelism(50)
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            loc.Lock()
            fib(10)
            loc.Unlock()
        }
    })
}

// 测试串行情况下，临界区为求斐波那契数列时，互斥锁性能
func BenchmarkMutexFib(b *testing.B) {
    loc := sync.Mutex{}
    for i := 0; i < b.N; i++ {
        loc.Lock()
        fib(10)
        loc.Unlock()
    }
}

// 测试并行情况下，临界区为求斐波那契数列时，互斥锁性能
func BenchmarkParallelMutexFib(b *testing.B) {
    loc := sync.Mutex{}
    b.SetParallelism(50)
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            loc.Lock()
            fib(10)
            loc.Unlock()
        }
    })
}


