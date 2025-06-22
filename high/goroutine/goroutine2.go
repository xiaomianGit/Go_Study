/**
* 指针，修改引用地址值
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	schedule := NewTaskSchedule()

	schedule.addStask("数据备份", func() {
		time.Sleep(2 * time.Second)
	})

	schedule.addStask("数据查询", func() {
		time.Sleep(3 * time.Second)
	})

	schedule.run()
	schedule.wg.Wait()
	schedule.PrintStats()
}

type Task struct {
	name      string
	execute   func()
	duration  time.Duration
	completed bool
}

type TaskSchedule struct {
	tasks []*Task
	wg    sync.WaitGroup
}

func (s *TaskSchedule) addStask(name string, execute func()) {
	s.tasks = append(s.tasks, &Task{
		name:    name,
		execute: execute})
}
func (s *TaskSchedule) run() {
	for _, task := range s.tasks {
		s.wg.Add(1)
		go s.executeTask(task)
	}
}
func (s *TaskSchedule) executeTask(task *Task) {
	defer s.wg.Done()

	start := time.Now()
	task.execute()
	task.duration = time.Since((start))
	task.completed = true

	fmt.Printf("任务 %s 执行完成耗时:%s \n", task.name, task.duration)
}
func (s *TaskSchedule) PrintStats() {
	fmt.Println("任务执行统计")
	for _, task := range s.tasks {
		status := "已完成"
		if !task.completed {
			status = "未完成"
		}
		fmt.Printf("任务:%s 状态%s 耗时%s\n", task.name, status, task.duration)

	}
}

func NewTaskSchedule() *TaskSchedule {
	return &TaskSchedule{tasks: make([]*Task, 0)}
}
