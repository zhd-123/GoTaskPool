package task

import "fmt"

//定义一个协程池
type Pool struct{
	//任务channel
	JobsChannel chan *Task
	//最大worker数量
	maxWorker int
}

func NewPool(num int) *Pool{
	p := Pool{
		JobsChannel:make(chan *Task),
		maxWorker:num,
	}
	return &p;
}
//协程池的任务worker
func (p *Pool)worker(id int){
	//永久的从jobsChannel取任务执行
	for task := range p.JobsChannel{
		task.Execute()
		fmt.Println("worker :",id,"执行了一个任务")
	}
}
//启动协程池
func (p *Pool)Run(){
	for i:=0;i<p.maxWorker;i++{
		//每个worker都是一个goroutine
		go p.worker(i)
	}
}

