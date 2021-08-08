package startup

import (
	"context"
	"github.com/JungleMC/java-edition/internal/config"
	"github.com/JungleMC/java-edition/internal/net"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
)

func Start(rdb *redis.Client) {
	PopulateDummyData(rdb)

	config.Get = &config.Config{}
	if err := env.Parse(config.Get); err != nil {
		panic(err)
	}

	_, err := net.Bootstrap(rdb, config.Get.ListenAddress, config.Get.ListenPort, config.Get.OnlineMode)
	if err != nil {
		panic(err)
	}
}

func PopulateDummyData(rdb *redis.Client) {
	favicon := "data:image/png;base64,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"

	rdb.Set(context.Background(), "config:description", "A JungleTree Server", 0)
	rdb.Set(context.Background(), "config:favicon", favicon, 0)
	rdb.Set(context.Background(), "config:max_players", 20, 0)
}
