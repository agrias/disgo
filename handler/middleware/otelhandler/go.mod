module github.com/disgoorg/disgo.handler/middleware/otelhandler

go 1.20

replace github.com/disgoorg/disgo => ../../../

require (
	github.com/disgoorg/disgo v0.16.6
	github.com/disgoorg/log v1.2.0
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/metric v1.16.0
	go.opentelemetry.io/otel/trace v1.16.0
)

require (
	github.com/disgoorg/json v1.1.0 // indirect
	github.com/disgoorg/snowflake/v2 v2.0.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/sasha-s/go-csync v0.0.0-20210812194225-61421b77c44b // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/exp v0.0.0-20220325121720-054d8573a5d8 // indirect
	golang.org/x/sys v0.1.0 // indirect
)
