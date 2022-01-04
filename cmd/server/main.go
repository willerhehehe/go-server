//go:generate swagger generate spec -o ../../internal/static/swagger.json
//go install github.com/go-swagger/go-swagger/cmd/swagger@latest

package main

import (
	"dcswitch/internal/config"
	httpApi "dcswitch/internal/inbound/http"
	c "dcswitch/pkg/config"
	"dcswitch/pkg/encrypt"
	"dcswitch/pkg/logger"
	"dcswitch/pkg/mysql"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
)

var (
	configPath = flag.String("config-path", "config.yaml", "ConfigFile Path")
)

func main() {
	// 1. init config options
	options := config.Options{}

	err := c.InitFromFile(*configPath, "yaml", &options)
	if err != nil {
		panic(fmt.Errorf("config init error: %v\n", err.Error()))
	}

	// 2. init logger
	logOpt := logger.Option{
		LogPath: options.Logger.LogPath,
	}
	logger.InitLogger(logOpt)

	// 3. init DB
	pwd, err := encrypt.Decrypt(options.DataBase.Password, options.SecKey)
	if err != nil {
		log.Fatalf("Encrypt error: %v\n", err)
	}
	o := mysql.DBConnOptions{
		User:     options.DataBase.Username,
		Pwd:      pwd,
		Host:     options.DataBase.Host,
		Port:     options.DataBase.Port,
		Database: options.DataBase.DBName,
	}
	mysql.DB.InitConn(o)

	// 4. init servers
	servers := make(map[string]*http.Server)
	servers["BaseServer"] = httpApi.InitServer()
	servers["PProfServer"] = httpApi.InitPProfServer()
	for n, s := range servers {
		go func(n string, s *http.Server) {
			log.Infof("%s listen on %v\n", n, s.Addr)
			err := s.ListenAndServe()
			switch err {
			case http.ErrServerClosed:
				log.Error(err)
			default:
				log.Fatal(err)
			}
		}(n, s)
	}
	// Graceful shutdown
	httpApi.GracefulStop(servers)
}
