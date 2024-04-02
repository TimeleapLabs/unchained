package scheduler

import (
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/scheduler/logs"
	"github.com/KenshiTech/unchained/scheduler/uniswap"
	"github.com/go-co-op/gocron/v2"
	"os"
	"time"
)

type Scheduler struct {
	scheduler gocron.Scheduler
}

type Task interface {
	Run()
}

func New(options ...func(s *Scheduler)) *Scheduler {
	s := &Scheduler{}

	var err error
	s.scheduler, err = gocron.NewScheduler()
	if err != nil {
		log.Logger.Error("Failed to create token scheduler.")
		os.Exit(1)
	}

	for _, o := range options {
		o(s)
	}

	return s
}

func WithEthLogs() func(s *Scheduler) {
	return func(s *Scheduler) {
		for name, duration := range config.App.Plugins.EthLog.Schedule {
			s.AddTask(duration, logs.New(name, config.App.Plugins.EthLog.Events))
		}
	}
}

func WithUniswapEvents() func(s *Scheduler) {
	return func(s *Scheduler) {
		for name, duration := range config.App.Plugins.Uniswap.Schedule {
			s.AddTask(duration, uniswap.New(name, config.App.Plugins.Uniswap.Tokens))
		}
	}
}

func (s *Scheduler) AddTask(duration time.Duration, task Task) {
	log.Logger.With("duration", duration).Info("Register a new task")

	_, err := s.scheduler.NewJob(
		gocron.DurationJob(duration),
		gocron.NewTask(task.Run),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)
	if err != nil {
		log.Logger.Error("Failed to schedule task.")
		os.Exit(1)
	}
}

func (s *Scheduler) Start() {
	s.scheduler.Start()
}
