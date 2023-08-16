/*
Copyright © 2023 zag13
*/
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/zag13/tophub/server/config"
)

var (
	envFile = flag.String("e", "../.env", "the env file")
	cfgFile = flag.String("c", "./etc/tophub-server.yaml", "the config file")
)

func init() {
	flag.Parse()

	if err := godotenv.Load(*envFile); err != nil {
		fmt.Println(fmt.Sprintf("Loading env file error: %v", err))
		fmt.Println("Using default config file: ./etc/tophub-cli.yaml")
	}

	viper.SetConfigFile(*cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("SERVER")

	if err := viper.Unmarshal(&config.C); err != nil {
		log.Fatalf("Unable to unmarshal config: %v", err)
	}
}

func main() {
	srv := &http.Server{
		Addr:    config.C.Port,
		Handler: setupRouter(),
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
