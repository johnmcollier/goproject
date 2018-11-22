FROM golang:latest 

# Install dep, a Go package manager
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy the project over to the container's Gopath
ADD . $GOPATH/src/goproject 
WORKDIR $GOPATH/src/goproject 

# Install dependencies
RUN dep ensure -v

RUN go build -o main . 
CMD ["/go/src/goproject/main"]
EXPOSE 8000
