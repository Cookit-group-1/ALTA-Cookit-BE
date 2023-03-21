FROM golang:1.19-alpine

RUN mkdir /ALTA-Cookit-BE

WORKDIR /ALTA-Cookit-BE

COPY ./ /ALTA-Cookit-BE

RUN go mod tidy

RUN go build -o alta-cookit-be

CMD ["./alta-cookit-be"]