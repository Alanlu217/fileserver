native:
    go build -o out/backend ./backend
    go build -o out/syne ./syne

build:
	docker build -t file .

run:
	docker run -it --rm --net="host" file

sh:
	docker run -it --rm --net="host" file zellij
