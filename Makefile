# Prefer tools we've installed locally to those elsewhere on the PATH.
export PATH := $(abspath .env/bin):$(PATH)

.PHONY: initvenv
initvenv:
	python3 -m venv .env
	.env/bin/pip install -r requirements.txt

	GOBIN='$(abspath .env/bin)' go install \
		google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN='$(abspath .env/bin)' go install \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: clean
clean:
	rm -rf .env
	find . -type d -name '__pycache__' | xargs rm -rf
	find certs -name '*.csr' -o -name '*.pem' | xargs rm -rf

.PHONY: certs
certs: ## Regenerate certificates and keys
	cd certs && \
		cfssl gencert -initca ca-csr.json | cfssljson -bare ca - && \
		cfssl gencert \
			-ca=ca.pem \
			-ca-key=ca-key.pem \
			-config=ca-config.json \
			-hostname=127.0.0.1 \
			-profile=server server-csr.json | cfssljson -bare server && \
		cfssl gencert \
			-ca=ca.pem \
			-ca-key=ca-key.pem \
			-config=ca-config.json \
			-profile=client client-csr.json | cfssljson -bare client

.PHONY: proto
proto: proto-py proto-go
proto: ## Regenerate code from protobuf definitions

# Why sed, see https://github.com/protocolbuffers/protobuf/issues/1491
.PHONY: proto-py
proto-py:
	rm -f client/api/api_pb2_grpc.py client/api/api_pb2.py
	.env/bin/python -m grpc_tools.protoc -I proto --proto_path=proto \
		--python_out=client/api --grpc_python_out=client/api proto/api.proto
	cd client/api && cat api_pb2_grpc.py | \
		sed -E 's/^(import api_pb2.*)/from client.api \1/g' > api_pb2_grpc.tmp && \
		mv -f api_pb2_grpc.tmp api_pb2_grpc.py

.PHONY: proto-go
proto-go:
	rm -f server/api/api.pb.go server/api/api_grpc.pb.go
	protoc -I. --go_out=. --go-grpc_out=. proto/api.proto

