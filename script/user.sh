export DATABASE_USER=users
export DATABASE_PASSWORD=Database_$DATABASE_USER
export DATABASE_NAME=$DATABASE_USER
GRPC_ADDR=0.0.0.0:9100 go run . user --dev $@
