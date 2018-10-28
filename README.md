# ~ smart monkey ~

`smart monkey` is your friend. It helps you to distribute nginx configurations to a high number of servers.

## What does it do?

`smart monkey` is a tool to deploy nginx configurations to multiple servers. Imagine you have to distribute 
different nginx configurations to different servers, test if their syntax is valid, 
reload nginx, log the whole process and compute some cool statistics?

## How does it work?

`smart monkey` runs on the server-side and it is similar to a `daemon`.
It listens to a queue named with the server hostname, process the message content to find out what configuration it needs to deploy/remove.
All configuration data is stored in a `storage backend`, serving as `source truth`.

## Running

To run this program, you first need `docker-compose` installed.

After installing it, you can run:

```
docker-compose up
```
It runs a docker container with a RabbitMQ and etcd ready to be used.

To start the program, execute the main handler:

```
go run main.go -lifetime 0s
```

This will make the program to execute forever or until the user interrupt the program.
