package channel

import (
	"fmt"
)

func TestChannel() {
	//管道联系
	mychannel := make(chan int, 12)
	mychannel <- 10
	customerOne := <-mychannel
	fmt.Printf("%v\n", customerOne)
}
