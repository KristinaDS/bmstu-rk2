package main

import (
	"flag"
	"github.com/ValeryBMSTU/web-rk2/internal/api"
	"github.com/ValeryBMSTU/web-rk2/internal/config"
	"github.com/ValeryBMSTU/web-rk2/internal/provider"
	"github.com/ValeryBMSTU/web-rk2/internal/usecase"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "./configs/example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(prv)
	srv := api.NewServer(cfg.IP, cfg.Port, use)

	srv.Run()
}
