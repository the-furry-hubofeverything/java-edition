module github.com/JungleMC/java-edition

go 1.16

require (
	github.com/JungleMC/sdk v0.0.0-20210809140359-e8dcfa68f6af
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/go-redis/redis/v8 v8.11.2
	github.com/google/uuid v1.1.2
	google.golang.org/protobuf v1.27.1
)

replace github.com/JungleMC/sdk => ../sdk
