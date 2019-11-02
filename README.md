# env
`
Starting docker containers for common stuffs: mysql, adminer, es, rabbitmq, ...
`

# CLI:
```bash
./cli -h
cli
your profile CLI
version: 1.0.1
usage: cli [options] [command [command's options]]

options:
        -p | --profile:    Profile that cli with take action on.

        -h | --help:       Help.

commands:
    docker:     Up and running dockers container
                All possible containers a listed in etc/docker
                Will create volume for corresponding containers in proc/<ENV>/docker

    profile:    Config profile.
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

    stop    Stop all docker containers.
            Not yet implement stop by name.

    clean   Clean up all docker containers within namespace (default is work_<ENV>).

            Options:

            [--cache]       To remove cache.
```
