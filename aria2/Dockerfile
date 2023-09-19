FROM alpine

RUN apk add --no-cache --update aria2
WORKDIR /private


ENV LISTEN_PORT=6888 \
	RPC_PORT=6800 \
	RPC_SECRET= \
	DISABLE_IPV6=true \
	ENABLE_DHT6=false

COPY ./aria2.conf ./aria2.conf
EXPOSE 6888
EXPOSE 6888/udp
CMD touch /aria2/aria2.session && \
	aria2c --conf-path ./aria2.conf \
	--dht-file-path=/dht/dht.data \
	--dht-file-path6=/dht/dht6.data \
	--save-session=/aria2/aria2.session \
	--input-file=/aria2/aria2.session \
	# --rpc-secret=$RPC_SECRET \
	--rpc-listen-port=$RPC_PORT \
	--listen-port=$LISTEN_PORT \
	--disable-ipv6=$DISABLE_IPV6 \
	--enable-dht6=$ENABLE_DHT6
