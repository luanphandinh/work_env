# env
`
Starting docker containers for common stuffs: mysql, adminer, es, rabbitmq, ...
`

# Start and running
examples:
* Docker
```bash
chmod +x ./cli

./cli docker run mysql adminer
```

* Help
```bash
./cli -h
cli
your env CLI
version: 1.0.0
usage: cli [options] [command [command's options]]

options:
    -e | --env :  Define environment that cli with take action on
                  Environment: Possible values ['dev', 'test']
                  Avoiding conflict data, accidentally delete dev data when running test

commands:
    docker    Up and running dockers container
              All possible containers a listed in etc/docker
              Will create volume for corresponding containers in proc/<ENV>/docker

    *         Help
```

* Docker services
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
