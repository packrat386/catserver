VERSION=`cat VERSION`

default: osx linux docker

catserver_linux_amd64: catserver.go VERSION
	GOOS=linux go build -ldflags='-s' -o catserver_linux_amd64

catserver_darwin_amd64: catserver.go VERSION
	GOOS=darwin go build -ldflags='-s' -o catserver_darwin_amd64

linux: catserver_linux_amd64
osx: catserver_darwin_amd64

docker: catserver_linux_amd64
	docker build --rm --tag "packrat386/catserver:$(VERSION)" .
	docker build --rm --tag "packrat386/catserver:latest" .

clean:
	rm -f catserver catserver_darwin_amd64 catserver_linux_amd64
