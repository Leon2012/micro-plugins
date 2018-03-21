## Micro api ratelimit plugin ##

### install  ###

    go get -u github.com/Leon2012/micro-plugins
	
### edit main.go ###

	`
	import (
		"github.com/Leon2012/micro-plugins/ratelimit"
		"github.com/micro/micro/api"
		"github.com/micro/micro/cmd"
	)
	
	func init() {
		api.Register(ratelimit.NewRatelimiter())
	}
	`

### build ###

    go build


### run ###

	`./micro api --limit_num=1`