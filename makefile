
build-chain-tip:
	go build -o ./bin/get_chain_tip examples/get_chain_tip/main.go
	install_name_tool -add_rpath /usr/local/lib ./bin/get_chain_tip

build-tx-show:
	go build -o ./bin/tx-show examples/tx-show/main.go
	install_name_tool -add_rpath /usr/local/lib ./bin/tx-show
