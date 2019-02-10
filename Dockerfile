FROM golang


RUN mkdir -p /go/src/github.com/angadsharma1016/nephron

ADD . /go/src/github.com/angadsharma1016/nephron

WORKDIR /go/src/github.com/angadsharma1016/nephron

# to install pdftotext for docconv
RUN apt-get update && apt-get install -y poppler-utils wv unrtf tidy && go get github.com/JalfResi/justext

RUN go get -t -v ./...

EXPOSE 3000

CMD ["go","run","main.go"]