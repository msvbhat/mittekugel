# Mittekugel

Experiments with Go, REST APIs and CI/CD with Travis CI


This is an experiment project to learn web/REST developement with Go. For
the completeness of the project, I have also implemented CI/CD with Travis CI.

## Description

The project aims to provide few APIs which can be used to metrics aggregation
and analysis. But at the moment only two APIs are implemented. This page will
be updated as and when I update them.

### GET /v1/metrics/node/{nodename}

At the moment this always returns 200 with name of the Node. This will be changed
to POST method soon.

### GET /v1/analytics/nodes/average/

Returns a sample JSON with below struct

{
    time_slice
    cpu_used
    mem_used
}

## Developement

Just `go get github.com/msvbhat/mittekugel` and start making changes.

### Dependencies

This project uses Gorilla project to serve APIs. To install the dependcies
after cloing this repo. Just run `make dep`. That installs all the dependency

## Continuos Integration

This project uses [Travis CI](https://travis-ci.org/msvbhat/mittekugel) for
CI/CD. At the moment each push triggers a Travis job which in turn runs a
lint tests and unit tests. Upon passing them, a docker image will be created
and pushed to docker hub. At this moment the Travis Job doesn't deploy the
application anywhere.

## Building and Running the Application

To build the application, you need to have go installed and dependencies
installed. You can build the application using the below command.
`make build`

That creates a binary called `mittekugel` which you can simply run. The
application always binds to port 8080.

### Running with docker

The project also has a Dockerfile which can be used to build a docker image.
But each push after undergoing few tests, creates a docker image and pushes
to https://hub.docker.com/r/msvbhat/mittekugel/. You can download docker image
with `docker pull msvbhat/mittekugel:${VERSION}`

And start the docker with the command

`docker run -d -p 8080:8080 msvbhat/mittekugel:${VERSION}`

## Deploying to Kubernetes

There is a [kubernetes deployment manifest](./deploy/k8s-deployment.yaml) with
this repo which can be used to deploy this application to a kubernetes cluster.

There is also a service manifest which exposes the deployment using the Kubernetes
NodePort type service. But in an real scenario some kind of ingress should
be used.

## Things to do

There are lot of things to do for this little side project. I hope to complete them
in the coming days. I have listed few of them here.

1. Implementing a CD from Travis to staging environment.
1. Improve the code by using a logger.
1. Implementing the APIs properly.
1. Use some kind of time series Databases to store the data.
1. Automated functional tests after deploying to staging environment.
