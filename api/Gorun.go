package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GoRun(ctx *gin.Context)  {
	ch := make(chan interface{}, 1)
	
	go Cus(ch)

	v := 1
	//for  {
		for i := 0; i < 5; i++ {
			ch <- v
			fmt.Printf("生产... %v \n", v)
			//time.Sleep(time.Second * 2)
			v++
		}
		//time.Sleep(time.Second * 2)
	//}
}

func Cus(ch chan interface{})  {
	for v := range ch {
		fmt.Printf("消费... %v \n", v)
	}
}