test: install tests

install:
	chmod +x cli
	chmod -R +x bin/

tests:
	chmod -R +x test/
	test/config-profile
