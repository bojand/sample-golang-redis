# sample-golang-redis

Sample app using redis

```sh
PORT=4000 REDIS_URL=... ./sample-golang-redis
```

```sh
$ http POST 'http://localhost:4000/counter'
HTTP/1.1 200 OK
Date: Thu, 25 Mar 2021 13:47:38 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

3

$ http 'http://localhost:4000/counter'
HTTP/1.1 200 OK
Date: Thu, 25 Mar 2021 13:47:44 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

3
```
