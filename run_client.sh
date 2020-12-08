export GO111MODULE="auto"

export SERVER_ADDRESS="localhost"
export PORT="50051"

export CLIENT_TYPE="calculate" # client_tpye("hi", "calculate")

go run ex_client/main.go ex_client/calculate.go ex_client/hi.go