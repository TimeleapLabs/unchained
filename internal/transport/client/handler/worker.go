package handler

type worker struct {
}

func NewWorkerHandler() Handler {
	return &worker{}
}
