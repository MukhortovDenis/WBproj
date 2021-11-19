FROM golang:1.17.1

WORKDIR /WBproj/
ADD ./ /WBproj/

RUN go mod download

RUN go build -o WBproj .

CMD [ "./WBproj" ]