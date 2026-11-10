PACKAGES := ./...
BUILD := build

all: s21_graph s21_graph_algorithms app

s21_graph:
	mkdir -p $(BUILD)
	go build -buildmode=archive -o ./$(BUILD)/s21_graph.a ./graph

s21_graph_algorithms:
	mkdir -p $(BUILD)
	go build -buildmode=archive -o ./$(BUILD)/s21_graph_algorithms.a ./graph_algorithms

app:
	mkdir -p $(BUILD)
	go build -o ./$(BUILD)/simplenavigator .

run:
	go run .

test:
	mkdir -p test_data/files_dot
	go test -race -cover $(PACKAGES)

gcov_report:
	mkdir -p test_data/files_dot
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

.PHONY: all s21_graph s21_graph_algorithms app run test gcov_report fmt clean
