build_docker:
	cd build && go run main.go

build_darwin:
	cd src && CGO_ENABLED=1 go build -buildmode=plugin -o ../bin/example.so .

install: build_darwin
	mkdir -p $(HOME)/.jumppad/plugins
	cp ./bin/example.so $(HOME)/.jumppad/plugins/example.so