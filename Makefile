test: install testing

install:
	chmod +x cli
	chmod -R +x bin/

testing:
	chmod -R +x test/
	test/config-profile
