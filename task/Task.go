package task

//定义任务封装
type Task struct{
	//可以定义任务数据字段

	fun func() error
}
//创建Task任务
func NewTask(arg func() error) *Task{
	t:=Task{
		fun:arg,
	}
	return &t
}
//执行Task任务
func (t *Task)Execute(){
	t.fun()
}
