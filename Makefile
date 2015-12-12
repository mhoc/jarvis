
build: jarvis

jarvis:
	gb build

deps:
	go get gopkg.in/yaml.v2
	go get github.com/gorilla/websocket
	go get gopkg.in/redis.v3

clean:
	rmdir bin pkg

run: build
	./bin/main

deploy:
	git pull
	gb build
	supervisorctl reload
