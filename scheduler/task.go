package scheduler

type Task struct{
    expression string
    command string
}

func NewTask(expression string, command string) Task{
    t := Task{expression,command}
    return t
}
