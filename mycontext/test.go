package main


import (
	"time"
	"fmt"
	"golang.org/x/net/context"
)

func Counter(ctx context.Context,compareChan chan int)  {
	count := 0
	go func(){
		time.Sleep(time.Second*6)
		fmt.Printf("\n%s----process---------------!", time.Now().Format("2006-01-02 15:04:05"))
		compareChan <- count
		return
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\n%s---Time out------------!", time.Now().Format("2006-01-02 15:04:05"))
			compareChan <- count
			return
		default:
		}
		count++
	}

}

func main() {
	t := time.NewTimer(time.Second)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		select {
		case <-t.C:
			compareChan := make(chan int)
			go Counter(ctx,compareChan)
			<-compareChan
			t.Reset(time.Second * 5)
		}
	}
}

//func main() {
//	ctx := context.Background()
//	ctx, cancel := context.WithCancel(ctx)
//	c := make(chan int)
//	go func() {
//		c <- Counter(ctx)
//	}()
//	time.Sleep(3 * time.Second)
//	cancel()
//	fmt.Println(<-c)
//
//}
