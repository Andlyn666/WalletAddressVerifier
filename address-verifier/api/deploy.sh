#!/bin/bash
# usage: ./deploy.sh ${path of the pem file} ${IP DNS of the EC2 instance}

PEMPATH=$1
EC2DNS=$2

# Build the excutable file
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o my_app .

# Kill old process
ssh -i $PEMPATH ec2-user@$EC2DNS "pkill my_app"

# Copy the excutable file to EC2 instance
scp -i $PEMPATH my_app ec2-user@$EC2DNS:~/.

# start the application
ssh -i $PEMPATH ec2-user@$EC2DNS "ls && chmod +x my_app && nohup ./my_app & "