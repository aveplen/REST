/*
    _____   _____   __  __   ____   __      _____   ______   _____
   /     \ /     \ |  \/  | |  _ \ |  |    |  ___| |_    _| |  ___|
   |  <--< |  |  | |      | |  __/ |  |__  |  ___|   |  |   |  ___|
   \_____/ \_____/ |_|\/|_| |__|   |_____| |_____|   |__|   |_____|

        _____   _____   _____   _____   _____   _____   _____
       /  ___\ /  _  \ |  _  \ |  __ \ /  _  \ /  ___\ |  ___|
       |  \_ \ |  _  | |     / |  __ < |  _  | |  \_ \ |  ___|
       \_____/ \_/ \_/ |__|__\ |_____/ \_/ \_/ \_____/ |_____|

*/

package main

import (
	"flag"

	"github.com/aveplen/REST/internal/config"
	"github.com/aveplen/REST/internal/database"
	"github.com/aveplen/REST/internal/handler"
	"github.com/aveplen/REST/internal/logger"
	"github.com/aveplen/REST/internal/server"
	"github.com/aveplen/REST/internal/service"
	"github.com/aveplen/REST/internal/store"
	_ "github.com/lib/pq"
)

var (
	configPath string
)

func main() {
	// parse configuration file path from flags
	flag.StringVar(
		&configPath,
		"config-path",
		"./config/config.yml",
		"path to configuration file",
	)
	flag.Parse()

	/*=================================================================*/
	/*                            Config                               */
	/*=================================================================*/
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	/*=================================================================*/
	/*                            Logger                               */
	/*=================================================================*/
	log, err := logger.NewLogger(cfg.Log)
	if err != nil {
		panic(err)
	}

	/*=================================================================*/
	/*                          Database                               */
	/*=================================================================*/
	db := database.NewDatabase(cfg.Pg)
	conn, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	if conn == nil {
		log.Fatal("fatal nahui")
	}
	defer db.Close()

	/*=================================================================*/
	/*                           Handler                               */
	/*=================================================================*/
	store := store.NewStore(conn)
	service := service.NewService(store, cfg.JWT)
	handler := handler.NewHandler(
		service,
		log,
	)

	/*=================================================================*/
	/*                            Server                               */
	/*=================================================================*/
	server := server.NewServer(handler, cfg.Srv)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
