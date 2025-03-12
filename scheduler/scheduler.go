package scheduler

import (
	"context"
	"order/dto"
	"order/service/orderservice"
	"sync"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type Config struct {
	SendOrdersToOutboxInterval int
}

type Scheduler struct {
	sch          gocron.Scheduler
	orderService orderservice.Service
	jobs         []gocron.Job
	config       Config
}

func New(config Config, orderService orderservice.Service) Scheduler {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	return Scheduler{
		config:       config,
		orderService: orderService,
		sch:          scheduler,
	}
}

func (s Scheduler) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	job, err := s.sch.NewJob(
		gocron.DurationJob(time.Duration(s.config.SendOrdersToOutboxInterval)*time.Millisecond),
		gocron.NewTask(s.PushOrdersToQueue),
	)
	if err != nil {
		panic(err)
	}

	s.jobs = append(s.jobs, job)

	s.sch.Start()

	<-ctx.Done()

	if err := s.sch.StopJobs(); err != nil {
		panic(err)
	}
}

func (s Scheduler) PushOrdersToQueue() {
	// TODO: tune the timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, err := s.orderService.PushToQueue(ctx, dto.PushToQueueRequest{})
	if err != nil {
		// TODO: log the error
	}
}
