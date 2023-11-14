protoc --go_out=. --go-grpc_out=. proto/db_service.proto 
protoc --dart_out=grpc:frontend/deputados/lib proto/db_service.proto

protoc --go_out=. --go-grpc_out=. proto/fotos_service.proto 
protoc --dart_out=grpc:frontend/deputados/lib proto/fotos_service.proto

protoc --go_out=. --go-grpc_out=. proto/server_service.proto 
protoc --dart_out=grpc:frontend/deputados/lib proto/server_service.proto
