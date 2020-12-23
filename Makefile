SHELL := /bin/bash
test: install tests

install:
	chmod +x cli.sh
	chmod -R +x bin/
	grep -q "alias cli=$(shell pwd)/cli.sh" ~/.bash_profile 2>/dev/null || echo "alias cli=$(shell pwd)/cli.sh" >> ~/.bash_profile
	grep -q "alias cli=$(shell pwd)/cli.sh" ~/.bashrc 2>/dev/null || echo "alias cli=$(shell pwd)/cli.sh" >> ~/.bashrc
	grep -q "alias cli=$(shell pwd)/cli.sh" ~/.zshrc 2>/dev/null || echo "alias cli=$(shell pwd)/cli.sh" >> ~/.zshrc

tests:
	chmod -R +x test/
	test/config-profile.sh

mac: nvim-install-mac nvim-config tmux-mac tmux-config
ubuntu: nvim-install-ubuntu nvim-config tmux-ubuntu tmux-config

nvim: nvim-install-mac nvim-config
nvim-ubuntu: nvim-install-ubuntu nvim-config
nvim-install-mac:
	brew install neovim
	brew install fd
	brew install ripgrep
	curl -fLo ~/.local/share/nvim/site/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

nvim-install-ubuntu:
	sudo apt-get install neovim
	sudo apt-get install fd-find
	sudo apt-get install ripgrep
	curl -fLo ~/.local/share/nvim/site/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

nvim-config:
	test -d ~/.config || mkdir ~/.config
	test -d ~/.config/nvim || mkdir ~/.config/nvim
	cp -r ./home/config/nvim/. ~/.config/nvim/
	nvim +PlugInstall +qall
	nvim -c 'CocInstall -sync|q'

tmux: tmux-mac tmux-config
tmux-mac:
	brew install tmux
	tmux new -d
	git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm
tmux-ubuntu:
	sudo apt-get update
	sudo apt-get install tmux
	tmux new -d
	git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm
tmux-config:
	cp ./home/tmux/.tmux.conf ~/.tmux.conf
	tmux source ~/.tmux.conf
	~/.tmux/plugins/tpm/scripts/install_plugins.sh

nodejs:
	curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.2/install.sh | bash
	nvm install 14
	nvm use 14

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

