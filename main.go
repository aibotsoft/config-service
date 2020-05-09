package main

import (
	"fmt"
	"github.com/aibotsoft/config-service/pkg/store"
	"github.com/aibotsoft/config-service/services/handler"
	"github.com/aibotsoft/config-service/services/server"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/aibotsoft/micro/mig"
	"github.com/aibotsoft/micro/sqlserver"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.New()
	log := logger.New()
	log.Infow("Begin service", "name", cfg.Service.Name, "config", cfg.ConfigService)
	db := sqlserver.MustConnectX(cfg)
	err := mig.MigrateUp(cfg, log, db)
	if err != nil {
		log.Fatal(err)
	}

	sto := store.NewStore(cfg, log, db)
	h := handler.NewHandler(cfg, log, sto)

	//sboApi:=api.NewApi(cfg, log )
	//au := auth.NewAuth(cfg, log, sto, sboApi)
	s := server.NewServer(cfg, log, h)

	// Инициализируем Close
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	//c:=collector.NewCollector(cfg, log, sboApi, sto, au)
	//go c.CollectJob()

	go func() { errc <- s.Serve() }()
	defer func() { s.Close() }()
	log.Info("exit: ", <-errc)
}
