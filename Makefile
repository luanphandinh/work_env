SHELL := /bin/bash
test: install tests

build:
	go build -o ./bin/build/cli -v ./src
	chmod +x ./bin/build/cli
	grep -q "alias cli2=$(shell pwd)/bin/build/cli" ~/.bash_profile 2>/dev/null || echo "alias cli2=$(shell pwd)/bin/build/cli" >> ~/.bash_profile
	grep -q "alias cli2=$(shell pwd)/bin/build/cli" ~/.bashrc 2>/dev/null || echo "alias cli2=$(shell pwd)/bin/build/cli" >> ~/.bashrc
	grep -q "alias cli2=$(shell pwd)/bin/build/cli" ~/.zshrc 2>/dev/null || echo "alias cli2=$(shell pwd)/bin/build/cli" >> ~/.zshrc

install:
	chmod +x cli.sh
	chmod -R +x bin/
	grep -q "alias cli=$(shell pwd)/cli.sh" ~/.bash_profile 2>/dev/null || echo "alias cli=$(shell pwd)/cli.sh" >> ~/.bash_profile
	grep -q "alias cli=$(shell pwd)/cli.sh" ~/.bashrc 2>/dev/null || echo "alias cli=$(shell pwd)/cli.sh" >> ~/.bashrc
	grep -q "alias cli=$(shell pwd)/cli.sh" ~/.zshrc 2>/dev/null || echo "alias cli=$(shell pwd)/cli.sh" >> ~/.zshrc

tests:
	chmod -R +x test/
	test/config-profile.sh
	go get ./...
	go test -v ./src/...
