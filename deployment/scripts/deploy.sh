set -e

SERVER=app@107.150.101.182
APPPATH="$GOPATH/src/github.com/iguiyu/demo/deployment"

cd $APPPATH

echo "Building application..."
CGO_ENABLED=0 GOOS=linux go build -o demo .

echo "Copy the executable file and script to the server"
scp demo $SERVER:~/server/tmp

rm demo

tar -cvf  $APPPATH/template.gz template
scp template.gz $SERVER:~/gopath/src/github.com/iguiyu/demo/deployment/tmp
rm -rf $APPPATH/template.gz


cd scripts
scp run_on_server.sh $SERVER:~/server/

echo "Run the script on the server"
ssh $SERVER -- /home/app/server/run_on_server.sh


# 1. tar -cvf  template.gz template
# 2. upload to server. gopath/src/github...
# 3. set gopath on run_on_server.sh
# 4. tar to unzip the static file to ... 