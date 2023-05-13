build:
	go build -o bin/ggcache

run: build
	./bin/ggcache --listenaddr :4000 --raftaddr :4001 --raftnode :9099/:3004/:5002 --localraftid node1&
	./bin/ggcache --listenaddr :4000 --raftaddr :4001 --raftnode :9099/:3004/:5002 --localraftid node1&


runfollower: build
	./bin/ggcache --listenaddr :4000 --leaderaddr :3000

test: 
	@go test -v ./...