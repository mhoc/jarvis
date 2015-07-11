
run: build
	./jarvis

build: jarvis

jarvis: deps main.go config
	go build github.com/mhoc/jarvis

deps:
	go get github.com/Sirupsen/logrus
	go get gopkg.in/yaml.v2
	go get golang.org/x/net/websocket

config: config/dynamic.go config/env.go config/yaml.go

clean:
	rm jarvis
