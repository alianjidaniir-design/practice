package main

import (
	"log/slog"
	"os"
	"time"
)

func myFunction() {
	j := 0
	for i := 1; i < 100000000; i++ {
		j = j*j%i + 9
	}
}
func main() {

	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)
	logger.Debug("This is a debug message")
	for i := 0; i < 5; i++ {
		now := time.Now()
		myFunction()
		elapsed := time.Since(now)
		logger.Info(
			"Observability",
			slog.Int64("elapsed", elapsed.Nanoseconds()),
		)
	}

}
