PACKAGES := ./...
BUILD := build ./test_data/files_dot/*

s21_graph:
	mkdir -p $(BUILD)
	go build -buildmode=archive -o ./$(BUILD)/s21_graph.a ./graph

test: 
	go test -race -cover $(PACKAGES)

gcov_report:
	go test $(PACKAGES) -coverprofile=coverage.txt
	go tool cover -html coverage.txt -o index.html
	open index.html

fmt:
#	Формотирование
	gofmt -l . 
# 	Потенциальные ошибки
	go vet $(PACKAGES) 
# 	Ревью кода
	golangci-lint run $(PACKAGES) 
# 	Актуальность mod
	go mod tidy -diff

clean:
	rm -rf $(BUILD) index.html coverage.txt 
