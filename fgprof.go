package pprof

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/felixge/fgprof"

	"github.com/hertz-contrib/pprof/adaptor"
)

const (
	// DefaultPrefix url prefix of pprof
	DefaultFpprofPrefix = "/debug/fgprof"
)

func getFgprofPrefix(prefixOptions ...string) string {
	prefix := DefaultFpprofPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}
	return prefix
}

func FgprofRegister(r *server.Hertz, prefixOptions ...string) {
	FgprofRouteRegister(&(r.RouterGroup), prefixOptions...)
}

func FgprofRouteRegister(rg *route.RouterGroup, prefixOptions ...string) {
	prefix := getFgprofPrefix(prefixOptions...)

	prefixRouter := rg.Group(prefix)
	{
		prefixRouter.GET("/", adaptor.NewHertzHTTPHandlerFunc(fgprof.Handler().ServeHTTP))
	}
}
