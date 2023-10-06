package main

import (
	"time"

	
)

type Pomodoro struct{
	Taskname string

	StartTime time.Time
}

var currentTask Pomodoro = Pomodoro{}


func AddTask(taskName string)Pomodoro{

	return Pomodoro{
		Taskname: taskName,
		StartTime: time.Now(),
	}	
}

func PrintTask(task Pomodoro)string{

	if (Pomodoro{}== task){
		return "The TaskList is empty."
	} else {
		return "Task - " + task.Taskname
	}
}
