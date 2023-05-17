build:
	go build -o bin/ggcache

run: build
	./bin/ggcache

runfollower: build
	./bin/ggcache --listenaddr :4000 --leaderaddr :3000

runfollower2: build
	./bin/ggcache --listenaddr :4001 --leaderaddr :3000

test: 
	@go test -v ./...