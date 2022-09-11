package crawler

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/anvh2/trading-boy/logger"
	"github.com/anvh2/trading-boy/models"
)

func TestCrawl(t *testing.T) {
	crawler := New(logger.NewDev(), &models.ExchangeConfig{}, []string{"BTCUSDT"})
	crawler.Start()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Server now listening")

	go func() {
		<-sigs
		// run hooks here
		close(done)
	}()

	fmt.Println("Ctrl-C to interrupt...")
	<-done
	fmt.Println("Exiting...")
}
