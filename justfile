native:
    go build -o out/mnemo ./mnemo
    go build -o out/syne ./syne

build:
	docker build -t file .

run: build
	docker run -it --rm --net="host" file

sh:
	docker run -it --net="host" --name="File" file zellij
