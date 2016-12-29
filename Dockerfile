FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/tracker /opt/bin/tracker
ADD ./conf /opt/conf

CMD ["/opt/bin/tracker -c /opt/conf/tracker.conf"]
