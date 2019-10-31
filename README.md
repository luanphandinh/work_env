# work_env
`
Starting docker containers for common stuffs: mysql, adminer, es, rabbitmq, ...
`

# Start and running
examples:
* Docker
```bash
chmod +x ./cli

./cli -e dev docker run mysql adminer
```

* Help
```bash
./cli -h

Options:
    -e | --env :  Define environment that cli with take action on
                  Environment: Possible values ['dev', 'test']
                  Avoiding conflict data, accidentally delete dev data when running test
Commands:
    docker    Up and running dockers container
              All possible containers a listed in etc/docker
              Will create volume for corresponding containers in proc/<ENV>/docker
    *         Help
```
