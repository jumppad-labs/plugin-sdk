build:
	go build -buildmode=plugin -o ./bin/example

install: build
	mkdir -p $(HOME)/.jumppad/plugins
	cp ./bin/example $(HOME)/.jumppad/plugins/example