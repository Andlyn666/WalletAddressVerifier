# WalletAddressVerifier
Verify a user indeed owns the private key to the wallet address

Usage of deploy.sh:
./deploy.sh ${path of the pem file} ${IP DNS of the EC2 instance}

You can try the API:
http://13.233.30.226:1323/get_message?address=96216849c49358b10257cb55b28ea603c874b05e
http://13.233.30.226:1323/verify?address=96216849c49358b10257cb55b28ea603c874b05&signedMessage=162deedf0fa57fa1ac69f40b79754bb737ba186f8eb34bfb39a9132da93d796c2a18d45eb0fd6c5479379ba3cd245076f8e1f9d34f3f62fa8e6cd560b1a498a201 