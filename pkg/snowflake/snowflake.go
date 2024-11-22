package snowflake

import (
	"sync"

	"github.com/aiagt/aiagt/apps/user/conf"
	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

func initSnowFlake() {
	n, err := snowflake.NewNode(int64(conf.Conf().Auth.SnowflakeNode))
	if err != nil {
		panic(err)
	}

	node = n
}

func Generate() snowflake.ID {
	once.Do(initSnowFlake)
	return node.Generate()
}
