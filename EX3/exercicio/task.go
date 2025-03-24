package exercicio

import "fmt"

type Task struct{
	name string
	next *Task
}

type TaskList struct{
	head *Task
}

func(l *TaskList) AddTask(name string){
	newNode := &Task{name: name}
	if l.head == nil{
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil{
		current = current.next
	}
	current.next = newNode
}

func (l *TaskList) CompleteTask(){
	if l.head != nil{
		l.head = l.head.next
	}
}


func (l *TaskList) ShowTask(){
	current := l.head
	for current != nil{
		fmt.Print(current.name, "\n")
		current = current.next
	}
}

