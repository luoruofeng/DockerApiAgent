# DockerApiAgent
DockerApi代理：Unix Socket转TCP Socket。


# 编译
```shell
GOOS=linux GOARCH=amd64 go build -o da .
```

# 启动
```shell
#启动basic
docker run . basic -c config.json

#docker swarm leave --force

#启动master 192.168.0.29是内网ip
docker run . master 192.168.0.29 -c config.json
#docker swarm join-token manager

#启动worker192.168.0.29是master的内网ip
go run . worker SWMTKN-1-0e9y07dj5g7yy6tis1mjr9kvwgh45yqtmk08gfp0juwrzpqg38-658m9xmpi62pbkzxv3xgh4x8k -r 192.168.0.29 -c config.json
```