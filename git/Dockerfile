FROM alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash man man-pages git git-doc openssh
ENTRYPOINT [ "/usr/bin/git" ]
