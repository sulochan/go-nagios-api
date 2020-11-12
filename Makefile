all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/linux/nagios_sidecar_api -a -tags netgo -ldflags '-w' .
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/nagios_sidecar_api -a -tags netgo -ldflags '-w' .
