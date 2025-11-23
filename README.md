# Gartisan By iqbalatma

Gartisan is command line tool that help you develop golang project. This project inspired by Laravel artisan.


## How to build
```shell
go build -o gartisan && sudo mv gartisan /usr/local/bin
```


## How to use 

### Create controller
```shell
gartisan make:controller user_controller 
```

### Create model
```shell
gartisan make:model user 
```

### Create service
```shell
gartisan make:service user_service 
```

### Create repository
```shell
gartisan make:service user_repository 
```