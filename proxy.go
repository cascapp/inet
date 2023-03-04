package cloud

import (
	"dcs/leadgen-srv/internal/config"
	"dcs/leadgen-srv/internal/version"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Proxy(c *gin.Context, newhost string, newpath string) {
	remote, err := url.Parse(newhost)
	if err != nil {
		fmt.Println(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = newpath
		req.URL.RawQuery = c.Request.URL.RawQuery
		req.URL.RawQuery = c.Request.URL.RawQuery
	}
	c.Writer.Header().Add(
		"X-Reverse-Proxy",
		"DCS GoServ v"+version.CurrentVersion+" "+config.Current().Environment,
	)
	proxy.ServeHTTP(c.Writer, c.Request)
}
