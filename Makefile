
MEMCACHED_CONTAINER_NAME=memcache
MEMCACHED_PORT=11211

build: 
	go build -o bin/app main.go 

run: 
	go run main.go

docker.down.memcached:
	docker container rm -f ${MEMCACHED_CONTAINER_NAME}

docker.up.memcached: 
	docker container run -d -p 11211:11211 --name ${MEMCACHED_CONTAINER_NAME} memcached 

docker.restart.memcached: docker.down.memcached docker.up.memcached