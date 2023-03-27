# How To Run

1. Clone : git clone git@github.com:ilfanikoinworks/GRPC_Client.git
2. Clear proto file
    - Client : make clean
3. Create proto file
    - Client : make gen
4. Running application
    - Client Add : go run client/main.go add {value deposit}
    - Client Get : go run client/main.go get
    - Server : go run server/main.go