package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"app-api/boot"
	"app-api/middlewares"
	"app-api/routes"

	"github.com/gin-gonic/gin"

	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "HTTP Web Server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := &boot.Config.App
		gin.SetMode(cfg.Mode)
		app := gin.Default()

		// 全局中间件
		app.Use(
			gin.Recovery(),
			middlewares.RequestID(),
			middlewares.Cors(),
			// middlewares.Logger(),
		)

		// 静态资源
		app.Static("/storege/files", "./storege/files")
		app.Static("/assets", "./public/assets")
		app.StaticFile("/favicon.ico", "./public/favicon.ico")

		// 注册路由
		routes.RegisterAPIRoutes(app)
		routes.RegisterWebRoutes(app)

		srv := &http.Server{
			Addr:    strings.Join([]string{cfg.Host, cfg.Port}, ":"),
			Handler: app,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		log.Println("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}

		log.Println("Server exiting")
	},
}
