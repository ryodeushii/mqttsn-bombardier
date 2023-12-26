install-globals:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
watch:
	air -c .air.toml
watch-server:
	air -c .air.server.toml