/*
Copyright © 2023 zag13
*/
package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/zag13/tophub/server/api/router"
	"github.com/zag13/tophub/server/bootstrap"
)

func main() {
	flag.Parse()

	app := bootstrap.App()

	srv := &http.Server{
		Addr:    app.Config.Port,
		Handler: router.Setup(app),
	}

	go func() {
		log.Printf("Listening and serving HTTP on %s\n", app.Config.Port)
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
