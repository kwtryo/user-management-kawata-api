package runner

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kwtryo/user-management-kawata-api/testutil"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	cfg := testutil.CreateConfigForTest(t)

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	eg.Go(func() error {
		return Run(ctx, router, cfg)
	})
	// サーバーの起動を待つため3秒休む
	time.Sleep(3 * time.Second)

	url := fmt.Sprintf("http://localhost:%d/ping", cfg.Port)
	t.Logf("try request to %q", url)

	rsp, err := http.Get(url)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	// サーバの終了動作を検証する
	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}

	// 戻り値を検証する
	want := "pong"
	assert.Equal(t, want, string(got))
}
