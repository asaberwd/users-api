SERVICE = users-api
COMMIT=`git rev-parse --short HEAD`

clean:
	rm -rf ./bin

test:
	go test ./...

build: clean
	#GOOS=linux GOARCH=amd64
	go build -v -a -tags users-api -o bin/$(SERVICE) -a --ldflags "-w \
	-X github.com/asaberwd/users-api" cmd/main.go

deploy-local: build
	sls deploy --stage local --verbose

deploy-test: build
	sls deploy --stage test --verbose
