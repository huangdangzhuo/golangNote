package main

import (
	"fmt"
	"sync"
	"time"
)

// 未加锁
// func add(count *int, wg *sync.WaitGroup)   {
// 	for i:= 0;i<1000;i++ {
// 		*count +=1
// 	}
// 	wg.Done()
// }
//
// func main()  {
// 	var wg sync.WaitGroup
// 	count := 0
// 	wg.Add(3)
// 	go add(&count,&wg)
// 	go add(&count,&wg)
// 	go add(&count,&wg)
//
// 	wg.Wait()
// 	fmt.Println("count的值为",count)
// }

// 枷锁
// func add(count *int, wg *sync.WaitGroup,lock *sync.Mutex)   {
// 	for i:= 0;i<1000;i++ {
// 		lock.Lock()
// 		*count +=1
// 		lock.Unlock()
// 	}
// 	wg.Done()
// }
//
// func main()  {
// 	var wg sync.WaitGroup
// 	lock := &sync.Mutex{}
// 	count := 0
// 	wg.Add(3)
// 	go add(&count,&wg,lock)
// 	go add(&count,&wg,lock)
// 	go add(&count,&wg,lock)
//
// 	wg.Wait()
// 	fmt.Println("count的值为",count)
// }

func main()  {
	lock := &sync.RWMutex{}
	lock.Lock()
	for i:=0;i<4;i++ {
		go func(i int) {
			fmt.Printf("第%d个协程准备开始...\n",i)
			lock.RLock()
			fmt.Printf("第%d个协程获取读锁,sleep 1s ,释放锁\n",i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}
	time.Sleep(time.Second*2)
	fmt.Println("准备释放写锁，读锁不再阻塞")
	// 写锁一释放，读锁就自由了
	lock.Unlock()

	// 由于会等到读锁全部释放，才能获得写锁
	// 因为这里一定会在上面 4 个协程全部完成才能往下走
	lock.Lock()
	fmt.Println("程序退出...")
	lock.Unlock()
}

