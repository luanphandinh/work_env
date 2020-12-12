SHELL := /bin/bash
test: install tests

build:
	go build -o ./bin/build/cli -v ./bin/src/cli
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
	go test -v ./bin/src/cli/...

sudo-user:
	test $(name)
	adduser $(name)
	usermod -aG sudo $(name)

nginx:
	test $(domain)
	@echo "Installing packages:"
	sudo apt update
	sudo apt install
	sudo apt install nginx
	@echo "Create new http static sites for $(domain):"
	sudo mkdir -p /var/www/$(domain)/html
	sudo chown -R ${USER}:${USER} /var/www/$(domain)/html
	sudo chmod -R 755 /var/www/$(domain)
	DOMAIN=$(domain) envsubst < ./etc/nginx/sites-available/index.template.html > /var/www/$(domain)/html/index.html
	@echo "Create new site for $(domain):"
	sudo chown -R ${USER}:${USER} /etc/nginx/sites-available/
	DOMAIN=$(domain) envsubst < ./etc/nginx/sites-available/template > /etc/nginx/sites-available/$(domain)
	sudo ln -sfn /etc/nginx/sites-available/$(domain) /etc/nginx/sites-enabled/
	sudo systemctl restart nginx
	@echo "Make $(domain) https:"
	sudo apt install certbot python3-certbot-nginx
	sudo ufw allow 'Nginx Full'
	sudo certbot --nginx -d $(domain) -d www.$(domain)
	sudo systemctl restart nginx

