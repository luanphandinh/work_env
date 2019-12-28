test: install tests

install:
	chmod +x cli
	chmod -R +x bin/

tests:
	chmod -R +x test/
	test/config-profile

homebrew-mac:
	/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

homebrew-linux:
	sh -c "$(curl -fsSL https://raw.githubusercontent.com/Linuxbrew/install/master/install.sh)"

nvim: nvim-install nvim-config
nvim-install:
	brew install neovim
	brew install fd
	curl -fLo ~/.local/share/nvim/site/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

nvim-config:
	test -d ~/.config/nvim || mkdir ~/.config/nvim
	cp ./home/config/nvim/init.vim ~/.config/nvim/
	cp ./home/config/nvim/coc-settings.json ~/.config/nvim/
	nvim +PlugInstall +qall

