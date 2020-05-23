package main

import (
	"TaskPool/task"
	"fmt"
	"time"
)

func main(){
	t := task.NewTask(func()error{
		fmt.Println(time.Now())
		return nil
	})

	p := task.NewPool(4)

	//把任务交给协程池
	go func(){
		for{
			p.JobsChannel <- t
		}
	}()
	p.Run()
	select {

	}
}
