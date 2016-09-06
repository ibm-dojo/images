all:
	docker build -t ibmdojo/debug -f Dockerfile .
	docker tag ibmdojo/debug debug

push:
	docker push ibmdojo/debug

clean:
	-docker rmi -f ibmdojo/debug debug
