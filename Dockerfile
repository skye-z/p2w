FROM chromedp/headless-shell
MAINTAINER skye-z <skai-zhang@hotmail.com>

RUN apt-get update -y \
    && apt-get install -y fonts-noto-cjk \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/

COPY p2w /usr/local/bin/

RUN chmod +x /usr/local/bin/p2w

EXPOSE 12800

ENTRYPOINT []
CMD [ "p2w","server" ]