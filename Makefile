app-config:
	cp example.config.yaml config.yaml

install_dependencies: mocker
	GIT_TERMINAL_PROMPT=1 go get ./...

mocker:
	GO111MODULE=on go install github.com/golang/mock/mockgen@v1.4.4