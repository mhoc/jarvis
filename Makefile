
run: build
	./jarvis

build: jarvis

jarvis: deps main.go
	go build github.com/mhoc/jarvis

deps:
	go get gopkg.in/yaml.v2
	go get github.com/gorilla/websocket
	go get github.com/xuyu/goredis

clean:
	rm jarvis
