package main

import (
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
	"net/http"
	"fmt"
	
	"github.com/go-kit/kit/log"

	"github.com/alex-ethier/crypto-manager/trade"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	// Logging domain.
	var logger log.Logger
	{
		w := log.NewSyncWriter(os.Stderr)
		logger = log.NewLogfmtLogger(w)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var svc trade.Service
	svc = trade.TradeService{}
	svc = trade.LoggingMiddleware(logger)(svc)
	endpoint := trade.Endpoints{
		TradeEndpoint: trade.MakeTradeEndpoint(svc),
	}

	r := trade.MakeHttpHandler(ctx, endpoint, logger)

	// HTTP transport
	go func() {
		logger.Log("msg", "Starting server at port 8080")
		handler := r
		errChan <- http.ListenAndServe(":8080", handler)
	}()


	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<- errChan)
}
