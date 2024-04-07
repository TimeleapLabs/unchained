package scheduler

import (
	"os"
	"time"

	"github.com/KenshiTech/unchained/internal/ethereum"

	"github.com/KenshiTech/unchained/internal/persistence"
	"github.com/KenshiTech/unchained/internal/scheduler/uniswap"
	evmLogService "github.com/KenshiTech/unchained/internal/service/evmlog"
	uniswapService "github.com/KenshiTech/unchained/internal/service/uniswap"

	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/scheduler/logs"
	"github.com/go-co-op/gocron/v2"
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

func WithEthLogs(
	evmLogService *evmLogService.Service,
	ethRPC *ethereum.Repository,
	persistence *persistence.BadgerRepository,
) func(s *Scheduler) {
	return func(s *Scheduler) {
		if config.App.Plugins.EthLog == nil {
			return
		}

		for name, duration := range config.App.Plugins.EthLog.Schedule {
			s.AddTask(duration, logs.New(
				name, config.App.Plugins.EthLog.Events,
				evmLogService, ethRPC, persistence,
			))
		}
	}
}

func WithUniswapEvents(
	uniswapService *uniswapService.Service,
	ethRPC *ethereum.Repository,
) func(s *Scheduler) {
	return func(s *Scheduler) {
		if config.App.Plugins.Uniswap == nil {
			return
		}

		for name, duration := range config.App.Plugins.Uniswap.Schedule {
			s.AddTask(duration, uniswap.New(
				name, config.App.Plugins.Uniswap.Tokens,
				uniswapService, ethRPC,
			))
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

	select {}
}
