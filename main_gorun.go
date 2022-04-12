package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan interface{})

	go Cus(ch)

	//for  {
	/*go func() {
		v := 1
		for i := 0; i < 5; i++ {
			ch <- v
			fmt.Printf("生产... %v \n", v)
			//time.Sleep(time.Second * 2)
			v++
		}
	}()*/
	/*for  {
		v := 1
		for i := 0; i < 5; i++ {
			ch <- v
			fmt.Printf("生产... %v \n", v)
			//time.Sleep(time.Second * 2)
			v++
		}
		time.Sleep(time.Second * 5)
	}*/
	//time.Sleep(time.Second * 2)
	//}
	//close(ch)

	/*for val := range ch {
		fmt.Println(val)
	}*/
	time.Sleep(time.Second * 10)
}

func Cus(ch chan interface{})  {
	for v := range ch {
		fmt.Printf("消费... %v \n", v)
	}
}