export DATABASE_USER=collection
export DATABASE_PASSWORD=Database_$DATABASE_USER
export DATABASE_NAME=$DATABASE_USER
GRPC_ADDR=0.0.0.0:9801 go run . collection --dev
