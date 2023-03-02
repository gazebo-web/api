build:
	@echo "Building gRPC and Protobuf in all languages"
	buf generate --template buf.gen.go.yaml
	buf generate --template buf.gen.ts.yaml
	buf generate --template buf.gen.cpp.yaml

clean:
	@echo "Removing compiled protobuf messages"
	rm -R */*.pb*
	rm -Rf cpp/
	rm -Rf ts/
