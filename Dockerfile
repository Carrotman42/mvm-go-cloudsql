FROM gcr.io/google_appengine/golang

COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install -tags appenginevm
