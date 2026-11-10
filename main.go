package main

import (
	"bufio"
	"fmt"
	"os"
	"simplenavigator/graph"
	graph_algorithms "simplenavigator/graph_algorithms"
	"strconv"
	"strings"
)

func printMenu() {
	fmt.Println("\n=== Меню ===")
	fmt.Println("1. Загрузить граф из файла")
	fmt.Println("2. Обход в ширину (BFS)")
	fmt.Println("3. Обход в глубину (DFS)")
	fmt.Println("4. Кратчайший путь между двумя вершинами")
	fmt.Println("5. Кратчайшие пути между всеми парами вершин")
	fmt.Println("6. Минимальное остовное дерево")
	fmt.Println("7. Задача коммивояжера")
	fmt.Println("0. Выход")
	fmt.Print("Ваш выбор: ")
}

func readLine(r *bufio.Reader) (string, error) {
	s, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(s), nil
}

func readInt(r *bufio.Reader, prompt string) (int, error) {
	fmt.Print(prompt)
	s, err := readLine(r)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var g *graph.Graph
	loaded := false

	for {
		printMenu()

		s, err := readLine(reader)
		if err != nil { // EOF (Ctrl+D) — выходим, иначе цикл станет бесконечным
			return
		}
		choice, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Введите число")
			continue
		}

		if choice >= 2 && choice <= 7 && !loaded {
			fmt.Println("Сначала загрузите граф (пункт 1)")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Введите путь до файла: ")
			filename, err := readLine(reader)
			if err != nil {
				return
			}
			if err := g.LoadGraphFromFile(filename); err != nil {
				fmt.Println("Ошибка загрузки:", err)
				continue
			}
			loaded = true
			fmt.Println("Граф загружен")

		case 2: // BFS
			start, err := readInt(reader, "Стартовая вершина: ")
			if err != nil {
				fmt.Println("Введите число")
				continue
			}
			res := graph_algorithms.BreadthFirstSearch(g, start) // ПРОВЕРЬТЕ: возвращает ли error
			//if err != nil {
			//	fmt.Println("Ошибка:", err)
			//	continue
			//}
			fmt.Println(res)

		case 3: // DFS
			start, err := readInt(reader, "Стартовая вершина: ")
			if err != nil {
				fmt.Println("Введите число")
				continue
			}
			res := graph_algorithms.DepthFirstSearch(g, start)
			fmt.Println(res)

		case 4: // кратчайший путь между двумя вершинами
			v1, err1 := readInt(reader, "Первая вершина: ")
			v2, err2 := readInt(reader, "Вторая вершина: ")
			if err1 != nil || err2 != nil {
				fmt.Println("Введите число")
				continue
			}
			dist, flag := graph_algorithms.GetShortestPathBetweenVertices(g, v1, v2)
			if !flag {
				fmt.Println("Ошибка:", err)
				continue
			}
			fmt.Println("Длина кратчайшего пути:", dist)

		case 5: // все пары
			matrix := graph_algorithms.GetShortestPathsBetweenAllVertices(g)
			//if err != nil {
			//	fmt.Println("Ошибка:", err)
			//	continue
			//}
			for _, row := range matrix {
				fmt.Println(row)
			}

		case 6: // остовное дерево
			tree, err := graph_algorithms.GetLeastSpanningTree(g)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			for _, row := range tree {
				fmt.Println(row)
			}

		case 7: // коммивояжёр
			res, err := graph_algorithms.SolveTravelingSalesmanProblem(g)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			fmt.Println("Маршрут:", res.Vertices)
			fmt.Println("Длина:", res.Distance)

		case 0:
			return

		default:
			fmt.Println("Нет такого пункта меню")
		}
	}
}
