FROM alpine
RUN apk update
RUN apk add curl docker
RUN mkdir /docker

ENV DOCKER_VERSION=1.12.0

RUN curl -s https://get.docker.com/builds/Linux/x86_64/docker-$DOCKER_VERSION.tgz | \
        tar -C /tmp/ -xvzf- && \
		mv /tmp/docker/* /docker/ && \
		rm -rf /tmp/docker

ENV PATH=/docker:$PATH

ENTRYPOINT [ "/docker/dockerd" ]
CMD [ "-H", "0.0.0.0:2375", "-H", "unix:///var/run/docker.sock" ]
