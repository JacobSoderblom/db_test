package main

import (
	"os"
	"github.com/go-kit/kit/log"
	"github.com/JacobSoderblom/db_test/pkg/inmem"
	"github.com/JacobSoderblom/db_test/pkg/user"
	"net/http"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	users := inmem.NewUserRepository()

	us := user.NewService(users)

	mux := http.NewServeMux()

	mux.Handle("/api/user/", user.MakeHandler(us, logger))

	http.Handle("/", mux)

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", "localhost:3333", "msg", "listening")
		errs <- http.ListenAndServe("localhost:3333", nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}