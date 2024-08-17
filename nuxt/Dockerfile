### syntax=docker/dockerfile:1.7-labs
FROM nginx AS main

COPY --chmod=644 nginx/conf.d /etc/nginx/conf.d
COPY --chmod=644 nginx/nginx.conf /etc/nginx/nginx.conf

RUN mkdir -p /var/log/nginx \
    && touch /var/log/nginx/upstream.log \
    && chmod 755 /var/log/nginx/upstream.log

#COPY .output/public /usr/share/nginx/html

FROM main

COPY app-go app-go

COPY .env .env

CMD ./app-go

EXPOSE 80
