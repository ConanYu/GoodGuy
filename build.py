import os
import shutil
import logging


def main() -> None:
    cwd = os.getcwd()
    root = os.path.dirname(os.path.abspath(__file__))
    os.chdir(root)
    os.system('protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative -I. CrawlService/crawl_service/crawl_service.proto')
    if not os.path.exists(os.path.join(root, 'src', 'pb')):
        os.mkdir(os.path.join(root, 'src', 'pb'))
    pb_src = os.path.join(root, 'CrawlService', 'crawl_service', 'crawl_service.pb.go')
    pb_dst = os.path.join(root, 'src', 'pb', 'crawl_service.pb.go')
    pb_grpc_src = os.path.join(root, 'CrawlService', 'crawl_service', 'crawl_service_grpc.pb.go')
    pb_grpc_dst = os.path.join(root, 'src', 'pb', 'crawl_service_grpc.pb.go')
    try:
        shutil.copyfile(pb_src, pb_dst)
    except Exception as e:
        logging.exception(e)
    try:
        shutil.copyfile(pb_grpc_src, pb_grpc_dst)
    except Exception as e:
        logging.exception(e)
    os.chdir(cwd)


if __name__ == '__main__':
    main()
