# env ![](https://github.com/luanphandinh/env/workflows/workspace/badge.svg) ![](https://github.com/luanphandinh/env/workflows/cli/badge.svg)
* Easy, clean and faster way to spin up docker containers: mysql, adminer, es, rabbitmq, ...
* Seprate profile and highly confirgurable.

# install
```bash
make install
source ~/.bash_profile
```

# test
```bash
make tests
```

# CLI:
```bash
./cli -h
```
```bash
cli.sh
your profile CLI
version: 1.1.1
usage: cli.sh [options] [command [command's options]]

options:
  -p | --profile <profile_name>:  Profile that cli with take action on.
                                  Auto create new one if not exist.
                                  defualt <profile_name>: default.

  -d | --debug:                   Turn on debug mode.

  -h | --help:                    Help.

commands:
  docker:           Up and running dockers container
                    All possible containers a listed in etc/docker
                    Will create volume for corresponding containers in var/lib/<profile_name>/docker

  set:              Config profile.
                    eg: ./cli.sh -p luanphan set SOME_VAR=SOME_VALUE OTHER_VAR=OVER_VALUE

  checkconf:        printenv of current profile to screen.
```

# Profile
* Configure your profile
```bash
./cli profile -n luanphan --set SOME_VAR=SOME_VALUE SOME_OTHER_VARS=SOME_OTHER_VALUE
```

* Help
```bash
profile
usage: profile [options]
options:
    -n | --name         Taking actions on profile.
                        If There is no profile, create one in etc/.
    -s | --set          Set ENV variables for profile.
    -h | --help         Help.
```

# Docker
* Up and running
Simply use `docker run` to start services that available with config from `etc/docker/*.yaml`. Use `default profile` if no profile are specified when running `cli`

```
./cli docker run all
```

* Help
```bash
cli docker -h

usage: docker [command]
commands:

    run     Start all specify docker containers.
            eg: docker run mysql adminer

            Options:

            [--like]    To run services that contain keyword
                        eg: docker run --like sql

    stop    Stop all docker containers within profile's name.(<default> is used if not specify).

    clean   Clean up all docker containers within profile's name (<default> is used if not specify).

            Options:

            [--cache]       To remove cache.
```

# Neovim
macOs
```
make nvim
```

Tips:
* Check the home/config/nvim/init.vim for key mapping.
* Increase keyboard setting to max `repeat key speed` to navigate faster.
* Could add gruvbox theme from home/theme for iterm.

# Tmux
macOs
```
make tmux
```

ubuntu:
```
make tmux-ubuntu
```
