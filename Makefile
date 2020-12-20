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
