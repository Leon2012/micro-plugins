package ratelimit

import (
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/micro/plugin"
	"go.uber.org/ratelimit"
)

const defaultLimitNum = 100

type ratelimiter struct {
	limiter ratelimit.Limiter
}

func (r *ratelimiter) Flags() []cli.Flag {
	return []cli.Flag{
		cli.IntFlag{
			Name:   "limit_num",
			Usage:  "limit number",
			EnvVar: "LIMIT_NUM",
		},
	}
}

func (r *ratelimiter) Commands() []cli.Command {
	return nil
}

func (r *ratelimiter) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, re *http.Request) {
			r.limiter.Take()
			h.ServeHTTP(rw, re)
		})
	}
}

func (r *ratelimiter) Init(ctx *cli.Context) error {
	limitNum := ctx.Int("limit_num")
	if limitNum <= 0 {
		limitNum = defaultLimitNum
	}
	log.Log("limit num : ", limitNum)
	r.limiter = ratelimit.New(limitNum)
	return nil
}

func (r *ratelimiter) String() string {
	return "Ratelimiter"
}

func NewRatelimiter() plugin.Plugin {
	r := &ratelimiter{}

	log.Log("call NewRatelimiter plugin")
	return r
}
