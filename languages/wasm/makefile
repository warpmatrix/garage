objdump-%: %.wasm
	wasm-objdump -d $<

xxd-%: %.wasm
	xxd -c 4 $<

%.wasm: %.wat
	wat2wasm $^ -o $@

clean:
	rm ./*.wasm
