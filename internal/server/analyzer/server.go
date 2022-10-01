package analyzer

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/anvh2/trading-bot/internal/logger"
	"github.com/anvh2/trading-bot/internal/pubsub"
	rpc "github.com/anvh2/trading-bot/internal/rpc/client"
	"github.com/anvh2/trading-bot/internal/storage"
	"github.com/anvh2/trading-bot/internal/worker"
	"github.com/anvh2/trading-bot/pkg/api/v1/notifier"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Server struct {
	logger    *logger.Logger
	notifyDb  *storage.Notify
	subcriber pubsub.Subscriber
	publisher pubsub.Publisher
	worker    *worker.Worker
	notifier  notifier.NotifierServiceClient
}

func New() *Server {
	logger, err := logger.New(viper.GetString("analyzer.log_path"))
	if err != nil {
		log.Fatal("failed to init logger", err)
	}

	redisCli := redis.NewClient(&redis.Options{
		Addr:       viper.GetString("redis.addr"),
		DB:         1,
		MaxRetries: 5,
	})

	publisher := pubsub.New(logger, redisCli)
	subscriber := pubsub.New(logger, redisCli)
	notifyDb := storage.NewNotify(logger, redisCli)

	worker, err := worker.New(logger, &worker.PoolConfig{NumProcess: 64})
	if err != nil {
		log.Fatal("failed to new workder", zap.Error(err))
	}

	conn, err := rpc.NewClient(viper.GetString("analyzer.notifier"), rpc.WithInsecure(), rpc.WithBlock())
	if err != nil {
		log.Fatal("failed to init notifier client conn", zap.Error(err))
	}

	notifier := notifier.NewNotifierServiceClient(conn)

	return &Server{
		logger:    logger,
		notifyDb:  notifyDb,
		subcriber: subscriber,
		publisher: publisher,
		worker:    worker,
		notifier:  notifier,
	}
}

func (s *Server) Start() error {
	s.worker.WithProcess(s.Process)

	s.subcriber.Subscribe(
		context.Background(),
		"trading.channel.analyze",
		func(ctx context.Context, message interface{}) error {
			s.worker.SendJob(ctx, message)
			return nil
		},
	)

	s.worker.Start()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Server now listening")

	go func() {
		<-sigs
		s.notifyDb.Close()
		s.publisher.Close()
		s.subcriber.Close()
		s.worker.Stop()

		close(done)
	}()

	fmt.Println("Ctrl-C to interrupt...")
	<-done
	fmt.Println("Exiting...")

	return nil
}