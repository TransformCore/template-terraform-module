.PHONY: test

test:
	(cd ./test/src && go test -v ./ -timeout 10m)

