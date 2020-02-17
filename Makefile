test: install tests

install:
	chmod +x cli.sh
	chmod -R +x bin/
	grep -q "alias cli=$(shell pwd)/cli.sh" ~/.bash_profile || echo "alias cli=$(shell pwd)/cli.sh" >> ~/.bash_profile

tests:
	chmod -R +x test/
	test/config-profile.sh

nvim: nvim-install nvim-config
nvim-install:
	brew install neovim
	brew install fd
	brew install the_silver_searcher
	curl -fLo ~/.local/share/nvim/site/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

nvim-config:
	test -d ~/.config/nvim || mkdir ~/.config/nvim
	cp ./home/config/nvim/init.vim ~/.config/nvim/
	cp ./home/config/nvim/coc-settings.json ~/.config/nvim/
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
