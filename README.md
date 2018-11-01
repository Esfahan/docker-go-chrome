## docker-go-chrome
Goland with headless chrome on CentOS7(Docker)

- chromedriver_linux64 2.43
- google-chrome-stable 70.0.3538.77
- Golang 1.11

## DockerHub
- https://hub.docker.com/r/esfahan/go-chrome

## Related
- https://github.com/Esfahan/docker-python-chrome

## docker build

```
$ sudo docker build -t go-chrome:latest ./
```

## docker run

```
$ sudo docker run --name go-chrome-container \
                  -v $(pwd)/code:/code \
                  -itd go-chrome:latest /bin/bash
```

## Initialize
### Create prj root directory

```
$ mkdir -p code/src/prj/
```

### go get dep

```
$ sudo docker exec -it go-chrome-container /bin/bash -c 'go get -u github.com/golang/dep/cmd/dep'
```

### dep init

```
$ sudo docker exec -it go-chrome-container /bin/bash -c 'cd /code/src/prj/ && dep init'
```

## Usage
### dep ensure

```
$ sudo docker exec -it go-chrome-container /bin/bash -c 'cd /code/src/prj && dep ensure'
```

### Execute

```
$ sudo docker exec -it go-chrome-container /bin/bash -c 'cd /code/src/prj && go run main.go'
Starting ChromeDriver 2.43.600233 (523efee95e3d68b8719b3a1c83051aa63aa6b10d) on port 32988
Only local connections are allowed.
2018/11/01 08:34:38 Google
2018/11/01 08:34:38 https://www.google.co.jp/
```
