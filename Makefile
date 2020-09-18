# Environment Variables
CGO=0
app=deployed
version?="0.0.0"
ldflags="-X main.version=$(version)"

default: build

build: clean
	env CGO_ENABLED=$(CGO) go build -ldflags $(ldflags) -o $(app) .

clean:
	rm -f $(app)

install:
	mv $(app) $(GOPATH)/bin/.	

