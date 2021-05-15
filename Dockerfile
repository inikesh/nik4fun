FROM golang:1.16.2
RUN mkdir /firstdoc
ADD . /firstdoc
WORKDIR /firstdoc
RUN go build -o local .
CMD ["/firstdoc/local"]