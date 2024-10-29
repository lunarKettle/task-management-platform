module auth_service

go 1.23.1

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.35.1
)

require github.com/joho/godotenv v1.5.1 // direct

require (
	golang.org/x/crypto v0.28.0 // direct
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
)
