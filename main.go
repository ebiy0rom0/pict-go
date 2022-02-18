package main

import (
	"context"
	"log"
	"net/http"
	"pict-go/gorilla"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

// エントリポイント
func main() {
	// context の準備 -> キャンセル処理の伝搬
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// gin(cors) 設定
	engine := gin.Default()
	// engine.Use(cors.New(cors.Config{
	// 	AllowWebSockets: true,
	// }))
	engine.Use(favicon.New("./favicon.ico"))

	// 実行
	run(ctx, engine)
}

func run(ctx context.Context, engine *gin.Engine) {
	// Hub 起動
	gorilla.RunHub(ctx)

	// test page
	engine.LoadHTMLFiles("index.html")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// websocket
	engine.GET("ws/", func(c *gin.Context) {
		gorilla.ServeWS(c.Writer, c.Request)
	})

	// web 起動
	err := engine.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
}
