package jobschedule

import (
	"container/heap"
	"fmt"
)

type Job struct {
	Name     string
	Priority int
	index    int
}

type jobQueue []*Job

type schedule struct {
	queue jobQueue
}

func (jq jobQueue) Len() int {
	return len(jq)
}

func (jq jobQueue) Less(i, j int) bool {
	return jq[i].Priority > jq[j].Priority
}

func (jq jobQueue) Swap(i, j int) {
	jq[i], jq[j] = jq[j], jq[i]
	jq[i].index = i
	jq[j].index = j
}

func (jq *jobQueue) Pop() interface{} {
	old := *jq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*jq = old[0 : n-1]
	return item
}

func (jq *jobQueue) Push(x interface{}) {
	n := len(*jq)
	item := x.(*Job)
	item.index = n
	*jq = append(*jq, item)
}

func (s *schedule) AddJob(j *Job) {
	heap.Push(&s.queue, j)
}

func (s *schedule) GetJob() (*Job, error) {
	if len(s.queue) == 0 {
		return nil, fmt.Errorf("job is empty")
	}
	item := heap.Pop(&s.queue)
	j := item.(*Job)
	return j, nil
}

func New() *schedule {
	s := &schedule{
		queue: make(jobQueue, 0),
	}
	heap.Init(&s.queue)
	return s
}
