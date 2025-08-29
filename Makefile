wasm:
	mkdir -p target/web
	cd target/web;\
	env GOOS=js GOARCH=wasm go build -o jampie.wasm github.com/niljimeno/jampie;\
	cp $$(go env GOROOT)/lib/wasm/wasm_exec.js .

	cp scripts/index.html target/web/

run-wasm:
	cd target/web;\
	python -m http.server;
