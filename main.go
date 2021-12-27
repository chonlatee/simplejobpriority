package main

import (
	"log"

	"github.com/chonlatee/jobschedule/jobschedule"
)

func main() {
	j := jobschedule.New()

	j.AddJob(&jobschedule.Job{
		Name:     "first job",
		Priority: 10,
	})

	j.AddJob(&jobschedule.Job{
		Name:     "second job",
		Priority: 9,
	})

	j.AddJob(&jobschedule.Job{
		Name:     "thrid job",
		Priority: 8,
	})

	jj, _ := j.GetJob()
	log.Println(jj.Priority, jj.Name)

	jj, _ = j.GetJob()
	log.Println(jj.Priority, jj.Name)

	jj, _ = j.GetJob()
	log.Println(jj.Priority, jj.Name)

	jj, err := j.GetJob()
	if err != nil {
		log.Println("get job err: ", err.Error())
	}

}
