# DockerApiAgent
DockerApi代理：Unix Socket转TCP Socket。


# 编译
```shell
GOOS=linux GOARCH=amd64 go build -o da .


# docker
docker build -t da:latest .

touch /da.log

#使用默认的配置文件启动basic
docker run -p 8888:8888 -v /da.log:/da.log -v /var/run/docker.sock:/var/run/docker.sock -v somewhere/config.json:/etc/da/config.json -e APP_CONFIG=/etc/da/config.json da:latest

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


# 其他
```shell
#启动consul

docker volume create consul-data
docker volume create consul-config

docker run -v consul-data:/consul/data -v consul-config:/consul/config -d --name consul -p 8500:8500 -p 8600:8600/udp consul agent -server  -bootstrap-expect=1 -ui -client=0.0.0.0
```