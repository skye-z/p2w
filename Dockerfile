FROM alpine:3.18.2
MAINTAINER skye-z <skai-zhang@hotmail.com>

RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk upgrade -U -a && \
    apk add chromium chromium-chromedriver\
    && rm -fr /var/cache/* && \
    mkdir /var/cache/apk

COPY p2w /usr/local/bin/

RUN cd /usr/local/bin/ && \
    chmod +x /usr/local/bin/p2w

EXPOSE 12800

CMD [ "p2w","server" ]