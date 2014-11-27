package scheduler

type Task struct{
    Expression string
    Command string
}

func NewTask(Expression string, Command string) Task{
    t := Task{Expression,Command}
    return t
}

