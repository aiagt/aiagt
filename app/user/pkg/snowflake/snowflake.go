package snowflake

import (
	"github.com/aiagt/aiagt/app/user/conf"
	"github.com/bwmarrin/snowflake"
	"sync"
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
