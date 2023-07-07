# AWS DVA-C02 Quiz App

This repository is dedicated to a quiz app geared towards passing the AWS DVA-C02 exam.

This application is built upon Docker/[Docker-Compose](https://docs.docker.com/compose/).

This application features the following software stack.

| Section   | Stack                                                                                                   |
| --------- | ------------------------------------------------------------------------------------------------------- |
| Frontend  | [NextJS](https://nextjs.org/), [Material-UI](https://mui.com/), [TailwindCSS](https://tailwindcss.com/) |
| API Layer | Golang+[GorillaMUX](https://github.com/gorilla/mux)                                                     |
| Backend   | [MariaDB Server](https://github.com/MariaDB/server)                                                     |

# Build / Installation Process

### Docker Compose Setup

For basic instructions on how to install Docker-Compose, [please follow the official instructions at this link.](https://docs.docker.com/compose/install/).

This specific application was built and tested using the version of `docker-compose` that came with MacOS Ventura + [Brew](https://brew.sh/). The commands below can be used to install docker+docker-compose.

```bash
brew install docker
brew install docker-compose
```

This application is designed to work with Docker-Compose.

The files found inside of `./Dockerfiles` contains the Dockerfile syntax used to create the Docker images `dva-ui` & `dva-api`. You will need to build these images before attempting to run this application.

Use the following instructions below if you need to build these images. Please rename the images to `Dockerfile` for each file found in `./Dockerfiles`.

```bash
docker build -t dva-ui ./Dockerfile
```
