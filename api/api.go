package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/krau/shisoimg/config"
	"github.com/krau/shisoimg/utils"
	"github.com/spf13/cobra"
)

func Serve(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	addr := config.Host()
	cmdHost := cmd.Flag("host")
	if cmdHost != nil {
		addr = cmdHost.Value.String()
	}
	utils.L.Infof("Listening on %s", addr)

	registerRoutes(r)

	if err := r.Run(addr); err != nil {
		utils.L.Fatal(err)
		os.Exit(1)
	}
}

func registerRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/random", randomImage)
	r.GET("/images/:md5", getImage)

	v1 := r.Group("/v1")
	artworkGroup := v1.Group("/artwork")
	{
		artworkGroup.Match([]string{http.MethodGet, http.MethodPost}, "/random", v1RandomArtworks)
		artworkGroup.Match([]string{http.MethodGet, http.MethodPost}, "/random/preview", randomImage)
		artworkGroup.Match([]string{http.MethodGet, http.MethodPost}, "/list", v1ListArtworks)
	}
}
