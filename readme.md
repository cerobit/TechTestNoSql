Second part of test

On this test i will use redis to track a voting system for favorite birds
Using redis

Local Install

* docker run --name redis-db -p 6379:6379 -d redis
* Start first microservice on port 8080

Test 

Using an http client test microservice with a post body that had a valid  the Bird id

### Send POST request with json body
POST http://localhost:8087/api/v1/voteup
Content-Type: application/json

{
"id": 3
}

The response 

HTTP/1.1 200 OK
Date: Mon, 20 Sep 2021 06:06:15 GMT
Content-Length: 54
Content-Type: text/plain; charset=utf-8

{"bird_id":3,"votes":9,"description":"Vote Complete"}


Response code: 200 (OK); Time: 37ms; Content length: 54 bytes