all:
	cd dind ; make all
	cd httpd ; make all
	cd looper ; make all
	cd tiny ; make all
	cd git ; make all

push:
	cd dind ; make push
	cd httpd ; make push
	cd looper ; make push
	cd tiny ; make push
	cd git ; make push

clean:
	cd dind ; make clean
	cd httpd ; make clean
	cd looper ; make clean
	cd tiny ; make clean
	cd git ; make clean
