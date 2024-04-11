GOOS=linux GOARCH=amd64 go build -o fuck
nohup ./fuck > nohup.log 2>&1 &
