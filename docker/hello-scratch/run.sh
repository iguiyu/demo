CGO_ENABLED=0 GOOS=linux go build -o main .
docker build -t kobeld/hello-scratch .
