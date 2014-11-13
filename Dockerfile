# build with:
# sudo docker build -t gcalc .

# execute with:
# sudo docker run -a stdin -a stdout -i -t --rm gcalc

FROM ubuntu

MAINTAINER Ryan Sheehan, rsheehan@gmail.com

RUN apt-get update
RUN apt-get install -y golang

RUN mkdir -p /go/src/github.com/AnarchyLime/gcalc

ADD . /go/src/github.com/AnarchyLime/gcalc

ENV GOPATH /go

CMD ["go", "run", "/go/src/github.com/AnarchyLime/gcalc/main.go"]
