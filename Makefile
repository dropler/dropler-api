.PHONY: docker-build docker-run docker gin
PORT=8080

docker-run:
	@docker run -t -e "PORT=$(PORT)" --publish 6060:$(PORT) --name droppio --rm droppio

docker-build:
	@docker build -t droppio .

docker: docker-build docker-run

gin:
	@gin -i
