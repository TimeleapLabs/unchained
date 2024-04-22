package scheduler

import (
	"os"
	"time"

	"github.com/KenshiTech/unchained/internal/scheduler/evmlog"

	"github.com/KenshiTech/unchained/internal/utils"

	"github.com/KenshiTech/unchained/internal/scheduler/uniswap"
	evmLogService "github.com/KenshiTech/unchained/internal/service/evmlog"
	uniswapService "github.com/KenshiTech/unchained/internal/service/uniswap"

	"github.com/KenshiTech/unchained/internal/config"
	"github.com/go-co-op/gocron/v2"
)

// Scheduler represents the scheduler service in the application and handles running tasks in a specific duration..
type Scheduler struct {
	scheduler gocron.Scheduler
}

// Task represents a task that can be scheduled by the scheduler.
type Task interface {
	Run()
}

// New creates a new scheduler service.
func New(options ...func(s *Scheduler)) *Scheduler {
	s := &Scheduler{}

	var err error
	s.scheduler, err = gocron.NewScheduler()
	if err != nil {
		utils.Logger.Error("Failed to create token scheduler.")
		os.Exit(1)
	}

	for _, o := range options {
		o(s)
	}

	return s
}

func WithEthLogs(
	evmLogService evmLogService.Service,
) func(s *Scheduler) {
	return func(s *Scheduler) {
		if config.App.Plugins.EthLog == nil {
			return
		}

		for name, duration := range config.App.Plugins.EthLog.Schedule {
			s.AddTask(duration, evmlog.New(
				name,
				evmLogService,
			))
		}
	}
}

func WithUniswapEvents(
	uniswapService uniswapService.Service,
) func(s *Scheduler) {
	return func(s *Scheduler) {
		if config.App.Plugins.Uniswap == nil {
			return
		}

		for name, duration := range config.App.Plugins.Uniswap.Schedule {
			s.AddTask(duration, uniswap.New(
				name,
				uniswapService,
			))
		}
	}
}

func (s *Scheduler) AddTask(duration time.Duration, task Task) {
	utils.Logger.
		With("duration", duration).
		Info("New UniSwap task scheduled")

	_, err := s.scheduler.NewJob(
		gocron.DurationJob(duration),
		gocron.NewTask(task.Run),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)
	if err != nil {
		utils.Logger.Error("Failed to schedule task")
		os.Exit(1)
	}
}

func (s *Scheduler) Start() {
	s.scheduler.Start()

	select {}
}
