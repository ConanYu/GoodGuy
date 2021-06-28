@echo off
cd %~dp0
cd ..
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative -I. CrawlService/crawl_service/crawl_service.proto
if not exist src\pb (
    cd src
    mkdir pb
    cd ..
)
copy CrawlService\crawl_service\crawl_service.pb.go src\pb\crawl_service.pb.go
copy CrawlService\crawl_service\crawl_service_grpc.pb.go src\pb\crawl_service_grpc.pb.go
del CrawlService\crawl_service\crawl_service.pb.go
del CrawlService\crawl_service\crawl_service_grpc.pb.go
cd %~dp0
