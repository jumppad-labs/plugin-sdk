build_plugin:
	cd build && go run main.go

install: build
	mkdir -p $(HOME)/.jumppad/plugins
	cp ./bin/example $(HOME)/.jumppad/plugins/example.so