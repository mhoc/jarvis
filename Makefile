
run: build
	./jarvis

build: jarvis

jarvis: deps main.go config.yaml config handlers log ws
	go build github.com/mhoc/jarvis

deps:
	go get gopkg.in/yaml.v2
	go get github.com/jbrukh/bayesian
	go get github.com/gorilla/websocket

config: config/env.go config/yaml.go

handlers: handlers/printer.go

log: log/console.go

ws: ws/reader.go ws/ws.go

clean:
	rm jarvis
