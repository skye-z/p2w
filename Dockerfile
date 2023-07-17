FROM alpine:3.18.2
MAINTAINER skye-z <skai-zhang@hotmail.com>

# 创建运行用户
RUN addgroup -S nonroot \
    && adduser -S betax -G nonroot

# 切换用户,不要使用root!!!
USER betax

# 如需在中国大陆地区构建清取消下方注释
# RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk upgrade -U -a && \
    apk add chromium chromium-chromedriver\
    && rm -fr /var/cache/* && \
    mkdir /var/cache/apk

COPY p2w /usr/local/bin/

RUN cd /usr/local/bin/ && \
    chmod +x /usr/local/bin/p2w

EXPOSE 12800

CMD [ "p2w","server" ]