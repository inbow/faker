.PHONY: generate_protobuf
generate_protobuf:
	@echo "Remove old generated *.pb.go"
	@rm *.pb.go
	@echo "Generate new *.pb.go"
	@protoc -I=. --go_out=. --go_opt=paths=source_relative *.proto