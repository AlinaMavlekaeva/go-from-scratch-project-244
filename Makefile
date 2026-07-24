build:
	go	build	-o	bin/gendiff	./cmd/gendiff
lint:
	golangci-lint	run
lint-fix:
	golangci-lint	run	--fix
test:
	go test ./code -cover
run1:
	./bin/gendiff ./testdata/fixture/file1.json ./testdata/fixture/file2.json -f plain
run2:
	./bin/gendiff ./testdata/fixture/filepath1.yml ./testdata/fixture/filepath2.yml -f plain
run3:
	./bin/gendiff ./testdata/fixture/filepath1.json ./testdata/fixture/filepath2.json -f stylish
run4:
	./bin/gendiff ./testdata/fixture/filepath1.json ./testdata/fixture/filepath2.json -f plain
run5:
	./bin/gendiff ./testdata/fixture/filepath1.json ./testdata/fixture/filepath2.json
