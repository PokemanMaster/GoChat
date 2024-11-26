#!/bin/bash

if ! [ -x "$(command -v docker-compose)" ]; then
  echo 'Error: docker-compose is not installed.' >&2
  exit 1
fi


http:// {
    root * /usr/share/caddy
    encode gzip
    file_server
}

lvyouwang.xyz {
    root * /usr/share/caddy
    encode gzip
    file_server
}

ee.lvyouwang.xyz {
    reverse_proxy 47.113.104.184:3000
    encode gzip
}

yna.lvyouwang.xyz {
    redir https://bilibili.com
}