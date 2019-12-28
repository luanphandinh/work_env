# env [![Build Status](https://travis-ci.org/luanphandinh/env.svg?branch=master)](https://travis-ci.org/luanphandinh/env)
* Easy, clean and faster way to spin up docker containers: mysql, adminer, es, rabbitmq, ...
* Seprate profile from each others.

# install
```bash
make install
```

# test
```bash
make test
```
* Note: `travis.yaml` also running some test to verify whether the docker services is bootstrap correctly.

# CLI:
```bash
./cli -h
```
```bash
cli
your profile CLI
version: 1.1.0
usage: cli [options] [command [command's options]]

options:
        -p | --profile:    Profile that cli with take action on.
        -d | --debug:      Turn on debug mode.

        -h | --help:       Help.

commands:
    docker:             Up and running dockers container
                        All possible containers a listed in etc/docker
                        Will create volume for corresponding containers in proc/<ENV>/docker

    config-profile:     Config profile.
    checkconf:          printenv of current profile to screen.
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
```
make nvim-mac
```
* Should check the home/config/nvim/init.vim for more information about key mapping to get the best experience
* Should increase keyboard setting to max `repeat key speed` to h/j/k/l move faster.
* Could use the gruvbox theme from home/theme for iterm.

