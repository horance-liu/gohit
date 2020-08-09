# gohit

## Build

```
$ git clone https://github.com/horance-liu/gohit.git && cd gohit
$ go build
```

## Docker Build

```
$ docker build -t gohit-dev:1.0 .
```

```
$ docker run -it -v `pwd`:/gohit gohit-dev:1.0 /bin/bash
```

```
$ docker-sync start
$ docker run -it -v gohit-sync:/gohit gohit-dev:1.0 /bin/bash
```
