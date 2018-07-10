FROM alpine:3.6

RUN apk --no-cache add \
    bash \
    openjdk8-jre \
    openssl \
    ruby \
    ruby-json \
    supervisor

RUN gem install --no-ri --no-rdoc sinatra

COPY etc/supervisor.d /etc/supervisor.d
COPY usr/local/bin /usr/local/bin
COPY usr/local/lib /usr/local/lib
#COPY usr/local/var /usr/local/var

CMD ["supervisord", "-n"]
