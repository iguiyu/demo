set -e

SERVER=app@107.150.101.182
APPPATH="$GOPATH/src/github.com/iguiyu/demo/deployment"

cd $APPPATH

echo "Building application..."
CGO_ENABLED=0 GOOS=linux go build -o demo .

echo "Copy the executable file and script to the server"
scp demo $SERVER:~/server/tmp

rm demo

cd scripts
scp run_on_server.sh $SERVER:~/server/

echo "Run the script on the server"
ssh $SERVER -- /home/app/server/run_on_server.sh
