# work_env
`
Starting docker containers for common stuffs: mysql, adminer, es, rabbitmq, ...
`

# Start and running
```bash
chmod +x run.sh
```
Options
```
-e  	Environment: Possible values ['dev', 'test']
    	Avoiding conflict data, accidentally delete dev data when running test
	Will create volume for corresponding containers

-s  	optional service to bootstrap container
```
Run all
```bash
./run.sh -e "dev" -s "all"
```
Or run optional services
```bash
./run.sh -e "dev" -s "mysql adminer es rabbitmq redis redis-commander"
```
