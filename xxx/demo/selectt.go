package demotest

import (
    "fmt"
    "time"
)

func SelectMain()  {
    bufChan := make(chan int)

    go func() {
        for{
            bufChan <-1
            time.Sleep(time.Second)
        }
    }()


    go func() {
        for{
            fmt.Println(<-bufChan)
        }
    }()

    // select{}
	for{}
}
