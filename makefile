
build-chain-tip:
	go build -o ./bin/get_chain_tip examples/get_chain_tip/main.go
	install_name_tool -add_rpath /usr/local/lib ./bin/get_chain_tip

build-chain-tip-and-run:
	go build -o ./bin/get_chain_tip examples/get_chain_tip/main.go
	install_name_tool -add_rpath /usr/local/lib ./bin/get_chain_tip
	./bin/get_chain_tip
