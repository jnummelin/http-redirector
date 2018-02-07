# HTTP Redirector

Simple service to redirect any HTTP request. It concatenates the requested path with a given base URL to be used as the redirect location. For example, with `http://bar.com` as the base URL:

`http://www.foo.com/foo/bar` get redirected to `http://bar.com/foo/bar`

## Build

`docker build -t jnummelin/http-redirector .`

## Running

Redirector picks up the base URL from env `REDIRECT_BASE`, so run with:

```
docker run -it --rm -e REDIRECT_BASE=https://foobar.com -p 8080:8080 jnummelin/http-redirector
```


### Docker compose

`docker-compose up --build`

Exposes port 8080 as ephemral port, so check `docker ps` list for the exposed port.