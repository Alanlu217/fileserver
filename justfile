build:
	docker build -t file .

run:
	docker run -it --rm -p8080:8080 file

sh:
	docker run -it --net="host" --name="File" file ash

