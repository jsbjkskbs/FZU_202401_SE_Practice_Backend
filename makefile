hz_new:
	hz new -module sfw 
	go mod tidy
	go mod edit -replace github.com/apache/thrift=github.com/apache/thrift@v0.13.0
	go mod tidy

hz_update:
	THRIFT_FILES="idl/user.thrift idl/video.thrift idl/relation.thrift idl/activity.thrift idl/interact.thrift idl/tool.thrift idl/report.thrift idl/oss.thrift"; \
	for file in $$THRIFT_FILES; do \
		hz update -module sfw -idl $$file; \
	done
	go mod tidy
	

run:
	go build
	./sfw

docker_build:
	docker build -t app:latest -f docker/Dockerfile .

docker_run:
	docker-compose -f docker/docker-compose.yml up -d