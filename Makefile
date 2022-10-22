.PHONY:
.SLIENT:

build: 
	go build -o ./.bin/bot cmd/bot/main.go
run: build
	./.bin/bot -config=config.local.yaml