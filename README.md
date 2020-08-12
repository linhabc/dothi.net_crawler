# Run

go run main.go data_type.go util.go db.go

# using docker

sudo docker run --name=crawler_go --mount source=output,destination=/app/output --mount source=db,destination=/app/db linhabc/dothi.net_crawler

# Using docker-compose

docker-compose up

# output folder

- output: store generated json file
- db: store generated leveldb folder
