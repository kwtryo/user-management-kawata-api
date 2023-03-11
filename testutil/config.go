package testutil

import (
	"log"
	"os"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/kwtryo/user-management-kawata-api/config"
)

// テスト用のconfigを返す
func CreateConfigForTest(t *testing.T) *config.Config {
	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		t.Fatalf("cannot parse env: %v", err)
	}
	cfg.Port = 8081
	// CI環境ならポート番号を変更
	if _, defined := os.LookupEnv("CI"); defined {
		log.Print("CI environment")
		cfg.DBPort = 5432
		// cfg.RedisPort = 6379
	}
	return cfg
}
