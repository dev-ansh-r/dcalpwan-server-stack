package redis

import (
	_ "embed"

	"github.com/redis/go-redis/v9"
)

var (
	//go:embed lua/dispatchTask.lua
	dispatchTaskScriptSource string
	dispatchTaskScript       = redis.NewScript(dispatchTaskScriptSource)
)
