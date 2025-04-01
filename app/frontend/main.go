// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by hertz generator.

package main

import (
	"context"
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzotelprovider "github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	rpc.InitClient()

	address := conf.GetConf().Hertz.Address

	p := hertzotelprovider.NewOpenTelemetryProvider(
		hertzotelprovider.WithEnableMetrics(false),
	)
	defer p.Shutdown(context.Background())

	h := server.New(server.WithHostPorts(address))
	h.LoadHTMLGlob("template/*")
	h.Delims("{{", "}}")
	registerMiddleware(h)

	router.GeneratedRegister(h)

	h.GET("sign-in", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-in", utils.H{
			"title": "Sign in",
			"next":  c.Query("next"),
		})
	})
	h.GET("sign-up", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-up", utils.H{
			"title": "Sign up",
		})
	})

	h.Static("/static", "./")

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}
	store, err := redis.NewStore(100, "tcp", conf.GetConf().Redis.Address, "", []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		panic(err)
	}
	store.Options(sessions.Options{MaxAge: 86400, Path: "/"})
	rs, err := redis.GetRedisStore(store)
	if err == nil {
		rs.SetSerializer(sessions.JSONSerializer{})
	}
	h.Use(sessions.New("cloudwego-shop", store))
	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())
	middleware.RegisterMiddleware(h)
}
