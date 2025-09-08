.PHONY: build
build:
	@docker build -t file .

.PHONY: run
run:
	@docker run -it --rm -p8080:8080 file

.PHONY: sh
sh:
	@docker run -it --rm -p8080:8080 file ash
