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

To run this program, you first need to run a docker image with a RabbitMQ up and running.

```
docker run -d --hostname smart-monkey --name smart-monkey -p 4369:4369 -p 5671:5671 -p 5672:5672 -p 15672:15672 rabbitmq:3
```

Now start the container:

```
docker start smart-monkey
```

It runs a docker container with a RabbitMQ, ready to be used.

You will want to run the administrative interface.

```
docker exec smart-monkey rabbitmq-plugins enable rabbitmq_management
```

To start the program, execute the main handler:

```
go run main.go -lifetime 0s
```

This will make the program to execute forever or until the user interrupt the program.
