# 自旋锁

为了对比自旋锁和互斥锁的性能，使用 Go 语言提供的 sync.Mutex 进行对比测试

当前机器硬件配置
处理器：2.6 GHz 六核Intel Core i7
内存：16G 

### 一、当临界区仅为 i++ 时 

#### 1. 不加锁且串行时，i++ 的耗时
执行命令：`go test -bench='BenchmarkAdd' -cpu=1,2,4,6,8,10,12`
```shell script
BenchmarkAdd            1000000000               0.254 ns/op
BenchmarkAdd-2          1000000000               0.255 ns/op
BenchmarkAdd-4          1000000000               0.274 ns/op
BenchmarkAdd-6          1000000000               0.276 ns/op
BenchmarkAdd-8          1000000000               0.276 ns/op
BenchmarkAdd-10         1000000000               0.255 ns/op
BenchmarkAdd-12         1000000000               0.255 ns/op
```

#### 2. 不加速且并行时，i++ 的耗时
执行命令：`go test -bench='BenchmarkParallelAdd' -cpu=1,2,4,6,8,10,12`
```shell script
BenchmarkParallelAdd            810551253                1.52 ns/op
BenchmarkParallelAdd-2          491213487                2.27 ns/op
BenchmarkParallelAdd-4          477727274                2.84 ns/op
BenchmarkParallelAdd-6          393606700                2.98 ns/op
BenchmarkParallelAdd-8          459784834                2.90 ns/op
BenchmarkParallelAdd-10         412693522                2.99 ns/op
BenchmarkParallelAdd-12         475581628                2.96 ns/op
```

#### 3. 使用自旋锁且串行时，i++ 的耗时
执行命令： `go test -bench='BenchmarkSpinLockAdd' -cpu=1,2,4,6,8,10,12` 
```shell script
BenchmarkSpinLockAdd            94733976                12.4 ns/op
BenchmarkSpinLockAdd-2          96934940                12.5 ns/op
BenchmarkSpinLockAdd-4          95506255                12.5 ns/op
BenchmarkSpinLockAdd-6          99826957                12.6 ns/op
BenchmarkSpinLockAdd-8          97101267                12.5 ns/op
BenchmarkSpinLockAdd-10         97458928                12.4 ns/op
BenchmarkSpinLockAdd-12         98872718                12.4 ns/op
```

#### 4. 使用自旋锁且并行时，i++ 的耗时
执行命令：`go test -bench='BenchmarkParallelSpinLockAdd' -cpu=1,2,4,6,8,10,12`
```shell script
BenchmarkParallelSpinLockAdd            93762098                13.0 ns/op
BenchmarkParallelSpinLockAdd-2          50967550                27.2 ns/op
BenchmarkParallelSpinLockAdd-4          42870268                30.1 ns/op
BenchmarkParallelSpinLockAdd-6          41822673                30.7 ns/op
BenchmarkParallelSpinLockAdd-8          40645779                28.4 ns/op
BenchmarkParallelSpinLockAdd-10         41570629                28.5 ns/op
BenchmarkParallelSpinLockAdd-12         42437306                28.6 ns/op
```

#### 5. 加互斥锁且串行时，i++ 的耗时
执行命令：`go test -bench='BenchmarkMutexAdd' -cpu=1,2,4,6,8,10,12`
```shell script
BenchmarkMutexAdd               100000000               11.1 ns/op
BenchmarkMutexAdd-2             100000000               11.1 ns/op
BenchmarkMutexAdd-4             100000000               11.4 ns/op
BenchmarkMutexAdd-6             100000000               11.1 ns/op
BenchmarkMutexAdd-8             100000000               11.2 ns/op
BenchmarkMutexAdd-10            100000000               11.0 ns/op
BenchmarkMutexAdd-12            100000000               11.0 ns/op
```

#### 6. 加互斥锁且并行时，i++ 的耗时
执行命令：`go test -bench='BenchmarkParallelMutexAdd' -cpu=1,2,4,6,8,10,12`
````shell script
BenchmarkParallelMutexAdd               56774598                24.1 ns/op
BenchmarkParallelMutexAdd-2             25215193                56.1 ns/op
BenchmarkParallelMutexAdd-4             15111068                96.1 ns/op
BenchmarkParallelMutexAdd-6             11604513               107 ns/op
BenchmarkParallelMutexAdd-8             12727336                96.0 ns/op
BenchmarkParallelMutexAdd-10            12935719                92.5 ns/op
BenchmarkParallelMutexAdd-12            13965248                89.2 ns/op
````


### 二、当临界区为求斐波那契数列

#### 1. 不加锁且串行时，斐波那契数列的耗时
执行命令：`go test -bench='BenchmarkFib' -cpu=1,2,4,6,8,10,12`
```shell script
BenchmarkFib             3966168               307 ns/op
BenchmarkFib-2           4040524               297 ns/op
BenchmarkFib-4           4189612               299 ns/op
BenchmarkFib-6           4018640               320 ns/op
BenchmarkFib-8           4166449               306 ns/op
BenchmarkFib-10          3785846               320 ns/op
BenchmarkFib-12          3485997               321 ns/op
```

#### 2. 不加锁且并行时，斐波那契数列的耗时
执行命令：`go test -bench='BenchmarkParallelFib' -cpu=1,2,4,6,8,10,12`
```shell script
BenchmarkParallelFib             4089374               300 ns/op
BenchmarkParallelFib-2           8380071               144 ns/op
BenchmarkParallelFib-4          16175791                72.8 ns/op
BenchmarkParallelFib-6          20471054                53.8 ns/op
BenchmarkParallelFib-8          25070583                47.8 ns/op
BenchmarkParallelFib-10         26416395                46.5 ns/op
BenchmarkParallelFib-12         27546380                50.6 ns/op
```

#### 3. 加自旋锁且串行时，斐波那契数列的耗时
执行命令：`go test -bench='BenchmarkSpinLockFib' -cpu=1,2,4,6,8,10,12`
```shell script
BenchmarkSpinLockFib             4036743               306 ns/op
BenchmarkSpinLockFib-2           3844212               312 ns/op
BenchmarkSpinLockFib-4           4021246               326 ns/op
BenchmarkSpinLockFib-6           3633670               320 ns/op
BenchmarkSpinLockFib-8           3379718               321 ns/op
BenchmarkSpinLockFib-10          3535051               354 ns/op
BenchmarkSpinLockFib-12          3493261               323 ns/op
```

#### 4. 加自旋锁且并行时，斐波那契数列的耗时
执行命令：`go test -bench='BenchmarkParallelSpinLockFib' -cpu=1,2,4,6,8,10,12` 
```shell script
BenchmarkParallelSpinLockFib             4039975               294 ns/op
BenchmarkParallelSpinLockFib-2           3625496               330 ns/op
BenchmarkParallelSpinLockFib-4           3555117               339 ns/op
BenchmarkParallelSpinLockFib-6           3372501               349 ns/op
BenchmarkParallelSpinLockFib-8           3079521               393 ns/op
BenchmarkParallelSpinLockFib-10          2930451               408 ns/op
BenchmarkParallelSpinLockFib-12          2895415               418 ns/op
```

#### 5. 加互斥锁且串行时，斐波那契数列的耗时
执行命令：`go test -bench='BenchmarkMutexFib' -cpu=1,2,4,6,8,10,12` 
```shell script
BenchmarkMutexFib                3743050               308 ns/op
BenchmarkMutexFib-2              4027713               311 ns/op
BenchmarkMutexFib-4              3925682               298 ns/op
BenchmarkMutexFib-6              3416304               319 ns/op
BenchmarkMutexFib-8              3653737               320 ns/op
BenchmarkMutexFib-10             4064810               329 ns/op
BenchmarkMutexFib-12             3756841               330 ns/op
```

#### 6. 加互斥锁且并行时，斐波那契数列的耗时
执行命令：`go test -bench='BenchmarkParallelMutexFib' -cpu=1,2,4,6,8,10,12`
```shell script
BenchmarkParallelMutexFib                3473006               320 ns/op
BenchmarkParallelMutexFib-2              3344409               361 ns/op
BenchmarkParallelMutexFib-4              3360722               357 ns/op
BenchmarkParallelMutexFib-6              3429223               364 ns/op
BenchmarkParallelMutexFib-8              3246004               358 ns/op
BenchmarkParallelMutexFib-10             3390898               350 ns/op
BenchmarkParallelMutexFib-12             3442885               350 ns/op
```