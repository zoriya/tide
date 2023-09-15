FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod ./
# go.sum
# RUN go mod download
COPY . .
RUN go build -o ./tide

EXPOSE 7890
CMD ./tide

