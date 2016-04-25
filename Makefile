
build: jarvis

jarvis:
	gb build

clean:
	rmdir bin pkg

run: build
	./bin/main

deploy:
	git pull
	gb build
	supervisorctl reload
