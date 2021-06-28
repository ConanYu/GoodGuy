package main

import (
	"GoodGuy/src/feishu"
	"GoodGuy/src/util"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"path"
	"syscall"
)

func loadViper() {
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	configPath := path.Dir(path.Dir(util.GetFileName()))
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func waitForShutdown() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}

func main() {
	loadViper()
	go feishu.Serve(viper.GetString("event.host"), viper.GetInt("event.port"))
	waitForShutdown()
}
