all:
	docker build -t ibmdojo/dind -f Dockerfile .
	docker tag ibmdojo/dind dind

push:
	docker push ibmdojo/dind

clean:
	-docker rmi -f ibmdojo/dind dind
