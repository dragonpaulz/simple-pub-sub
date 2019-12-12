## Requirements
Assumes golang, docker and docker-compose are installed on system. Docker and docker-compose are being used to get an instance of redis running. If there is a connection available to a redis instance running, then you may skip having docker and docker-compose installed.

I ran my application on Ubuntu 18.04 using: 
- go version go1.13.5 linux/amd64
- Docker version 18.09.7, build 2d0083d
- docker-compose version 1.17.1, build unknown

## To run
The first two steps are optional if you want to install redis on your local machine. If you have access to an instance of redis, you can change values in config.json file to connect to an already existing instance of redis.
`docker-compose build` to build redis image
`docker-compose up -d` to run redis

The next steps are to get a terminal window for each service. From the project's base directory, where the makefiles are located, run:
- `make publisher`
- `make sumfinder`
- `make meanfinder`

## Potential Improvements
- Get docker-compose to build images of publisher and subscribers, as well as get the wiring hooked up between images. The advantage of this approach is that setup and running the application would be a single command, and everyone running the application will have the same environment.
- Increase test coverage. Most of the code is using other libraries. I did however attempt to use an interface to break up some coupling to be able to stub out methods from the redis library so that I can insert my own tests. Due to time constraints, I have left it untouched in `feature/maketestable`.
- Adding flag to define the range of numbers generated by the publisher.
- Adding flag to define which configuration file to use, rather than expecting it to be the first argument after the program name.