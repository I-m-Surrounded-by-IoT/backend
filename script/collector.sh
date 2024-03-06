if [ ! "$GRPC_ADDR" ]; then
    GRPC_ADDR=0.0.0.0:11000
fi
go run . collector --dev
