FROM golang:1.15 as build
WORKDIR /go/src/github.com/webdevops/azure-scheduledevents-manager

# Get deps (cached)
COPY ./go.mod /go/src/github.com/webdevops/azure-scheduledevents-manager
COPY ./go.sum /go/src/github.com/webdevops/azure-scheduledevents-manager
COPY ./Makefile /go/src/github.com/webdevops/azure-scheduledevents-manager
RUN make dependencies

# Compile
COPY ./ /go/src/github.com/webdevops/azure-scheduledevents-manager
RUN make test
RUN make lint
RUN make build
RUN ./azure-scheduledevents-manager --help

#############################################
# FINAL IMAGE
#############################################
FROM ubuntu:20.04
ENV LOG_JSON=1
COPY --from=build /go/src/github.com/webdevops/azure-scheduledevents-manager/azure-scheduledevents-manager /
USER 1000
ENTRYPOINT ["/azure-scheduledevents-manager"]
