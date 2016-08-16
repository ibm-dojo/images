all:
	go build -tags netgo -installsuffix netgo .
	docker build -t httpd .
	docker tag httpd ibmdojo/httpd

push:
	docker push ibmdojo/httpd

clean:
	rm -f httpd
	-docker rmi -f httpd ibmdojo/httpd
