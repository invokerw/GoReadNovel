# Pull base image  
FROM centos:7

# FROM golang:1.6
#Gopath  /go

MAINTAINER my weifei "784856034@qq.com"

RUN mkdir /go
WORKDIR /go
RUN mkdir src
RUN mkdir bin
RUN mkdir pkg


# Add go path

# COPY ./github.com/ /root/gopro/src/
# COPY ./golang.org/ /root/gopro/src/
# COPY ./GoReadNovel/ /root/gopro/src/ 
# WORKDIR /root/gopro/src/GoReadNovel

# COPY ./github.com/ /go/src/ useless
# COPY ./golang.org/ /go/src/ useless
RUN yum install -y git
# WORKDIR /root/

RUN mkdir -p /go/src/github.com/golang/net/
WORKDIR /go/src/github.com/golang/net/

RUN git init
RUN git pull https://github.com/golang/net.git

RUN mkdir -p /go/src/golang.org/x/
RUN cp -R /go/src/github.com/golang/net /go/src/golang.org/x/


RUN yum install -y go
RUN echo "export GOPATH=/go">>/etc/profile
RUN echo "export PATH=$PATH:/go/bin">>/etc/profile
ENV GOPATH /go
ENV PATH $PATH:/go/bin

WORKDIR /go
RUN source /etc/profile

RUN go get github.com/astaxie/beego && go get github.com/beego/bee
RUN go get github.com/gin-gonic/gin 
RUN go get github.com/op/go-logging

EXPOSE 443

# ENTRYPOINT go run main.go
CMD ["bee", "run"]

