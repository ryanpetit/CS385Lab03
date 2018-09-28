FROM golang:alpine
COPY src/goApp.go src/
WORKDIR src/
CMD ["go", "run", "goApp.go"]
