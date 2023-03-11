package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kwtryo/user-management-kawata-api/clock"
	"github.com/kwtryo/user-management-kawata-api/config"
	"github.com/kwtryo/user-management-kawata-api/handler"
	"github.com/kwtryo/user-management-kawata-api/service"
	"github.com/kwtryo/user-management-kawata-api/store"
)

func SetupRouter(cfg *config.Config) (*gin.Engine, func(), error) {
	router := gin.Default()

	db, cleanup, err := store.New(cfg)
	if err != nil {
		return nil, cleanup, err
	}

	cl := clock.RealClocker{}
	r := &store.Repository{Clocker: cl}

	healthHandler := &handler.HealthHandler{
		Service: &service.HealthService{DB: db, Repo: r},
	}

	router.Use(CorsMiddleware())

	router.GET("/health", healthHandler.HealthCheck)

	return router, cleanup, nil
}

// CORSの設定
func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// preflight requestで許可した後の接続可能時間
		MaxAge: 24 * time.Hour,
	})
}
