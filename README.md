[![alt tag](https://drive.google.com/uc?export=view&id=16WXmLVKFNKxm_DXQMYCFi_A63A3YQdDD)](https://www.pngkey.com/detail/u2e6t4e6r5u2q8i1_microsoft-oss-conference-database-gopher-golang/)

[![License](https://img.shields.io/github/license/tedrong/topics)](https://github.com/tedrong/topics/blob/main/LICENSE) [![GitHub release (latest by date)](https://img.shields.io/github/v/release/tedrong/topics)](https://github.com/tedrong/topics/releases) [![Go Version](https://img.shields.io/github/go-mod/go-version/tedrong/topics)](https://github.com/tedrong/topics/blob/master/go.mod) [![DB Version](https://img.shields.io/badge/DB-PostgreSQL--latest-blue)](https://github.com/tedrong/topics/blob/master/go.mod) [![DB Version](https://img.shields.io/badge/DB-Redis--latest-blue)](https://github.com/tedrong/topics/blob/master/go.mod)

Welcome, trying to make your own **Topics**

This project contains complete service package example from backend service to user interface. We want to provide a infrastructure which can easily add services for any topics appeals to us. In this way, it can be used as a reference to generating information for users on web service.

The api server is base on [gin-boilerplate](https://github.com/Massad/gin-boilerplate) which already structured with [Gin Framework](https://gin-gonic.github.io/gin/), we add some new features such as json log system, renew database manipulate library to GORM, function scheduler system and selenium project to help us creating data crawler by web browser. Besides that, please check branch [www](https://github.com/tedrong/topics/tree/www#readme) for more detailed informations about our web application.
## Dependency

- [gorm](https://github.com/go-gorm/gorm): Full-Featured ORM for golang
- [zerolog](https://github.com/rs/zerolog): Write JSON format log
- [corn](https://github.com/robfig/cron): Job scheduler
- [Selenium](https://www.selenium.dev/): Create a browser-based data crawler
- API test

### Installation

```
$ go mod tidy
```

```
$ go install
```
##### Database Setup

We recommand installing both **postgresql** and **redis** on docker, it will reducing the task of installing and running software to as little as two commands (docker run and docker pull).

```
$ docker pull postgres
```

```
$ docker run --rm --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres
```

```
$ docker pull redis
```

```
$ docker run --rm -itd --name redis -p 6379:6379 redis
```

You will find the ***.sql** in `./database`, and you can import the postgres database using this command:

```
$ psql -U postgres -h localhost < ./database/xxx.sql
```
Note:
You can choose to create database by your own, when all databases have existed, gorm will create all tables automatically at server startup.

##### Selenium

Webdriver need additional dependencies, run the init.go in ./lib. The process will download all the files.

```
$ cd ./lib
$ go run init.go
```

Another requirement is Xvfb. According to official document, it is an X server that can run on machines with no display hardware and no physical input devices.  It emulates a dumb framebuffer using virtual memory.
Using apt command below, of you can comment out **selenium.StartFrameBuffer()** in ./crawler/entry.go.

```
$ sudo apt install -y xvfb
```

## Startup

Check .env and place your credentials

```
LOG_LEVEL: The logs which below this hierarchy will be drop.
SELENIUM: Absolute path of selenium .jar.
GECKO_DRIVER: Absolute path of mozilla geckodriver .jar.
```

> Make sure to change the values in .env for your databases

Generate SSL certificates (Optional)

> If you don't SSL now, change `SSL=TRUE` to `SSL=FALSE` in the `.env` file

```
$ mkdir cert/
```

```
$ sh generate-certificate.sh
```

```
$ go run *.go
```
## Testing

```
$ go test -v ./tests/*
```

## Import Postman Collection (API's)

Download [Postman](https://www.getpostman.com/) -> Import -> Import From Link

https://www.getpostman.com/collections/f485e28a80a8d6682016

Includes the following:

- User
  - Login
  - Register
  - Renew
  - Logout
- Auth
  - Refresh Token
- Dashboard
  - Info
  - Info_History
  - Client_Type_Percentage
  - Log

> In Login request in Tests tab:

```
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);

    var jsonData = JSON.parse(responseBody);
    pm.globals.set("token", jsonData.token.access_token);
    pm.globals.set("refresh_token", jsonData.token.refresh_token);

});
```

It captures the `access_token` from the success login in the **global variable** for later use in other requests.

Also, you will find in each request that needs to be authenticated you will have the following:

    Authorization -> Bearer Token with value of {{token}}

It's very useful when you want to test the APIs in Postman without copying and pasting the tokens.
