setup:
	go install github.com/jstemmer/go-junit-report/v2@latest

test-report:
	rm -rf artifacts
	mkdir artifacts
	go test -v 2>&1 ./... | go-junit-report -set-exit-code > artifacts/report.xml
