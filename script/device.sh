export DATABASE_USER=device
export DATABASE_PASSWORD=Database_$DATABASE_USER
export DATABASE_NAME=$DATABASE_USER
export GRPC_ADDR=0.0.0.0:10100
go run . device --dev
