package main

import (
	"fmt"
	"math"
)

// Структура для представления ребра
type Node struct {
	dest   string // Город
	weight int    // Вес (длина пути)
	next   *Node  // Следующий город
}

// Структура для графа
type Graph struct {
	adjList map[string]*Node // Список городов
}

// Функция для инициализации графа
func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[string]*Node),
	}
}

// Функция для добавления дороги между городами (ребро)
func (g *Graph) AddEdge(src, dest string, weight int) {
	node := &Node{
		dest:   dest,
		weight: weight,
		next:   g.adjList[src],
	}
	g.adjList[src] = node
}

// Алгоритм Дейкстры для нахождения кратчайших путей
func (g *Graph) Dijkstra(start string) map[string]int {
	dist := make(map[string]int) // Расстояние от начального города
	for city := range g.adjList {
		dist[city] = math.MaxInt // Инициализируем расстояния как бесконечность
	}
	dist[start] = 0

	visited := make(map[string]bool) // Множество посещённых городов

	// Основной цикл алгоритма
	for len(visited) < len(g.adjList) {
		// Находим город с минимальным расстоянием
		minDist := math.MaxInt
		var u string
		for city, d := range dist {
			if !visited[city] && d < minDist {
				minDist = d
				u = city
			}
		}

		// Если нет доступных городов для посещения, выходим
		if u == "" {
			break
		}

		visited[u] = true

		// Обновляем расстояния до соседей
		for node := g.adjList[u]; node != nil; node = node.next {
			if dist[u]+node.weight < dist[node.dest] {
				dist[node.dest] = dist[u] + node.weight
			}
		}
	}

	return dist
}

// Функция для вывода кратчайших путей от начального города
func (g *Graph) PrintShortestPaths(start string) {
	distances := g.Dijkstra(start)
	fmt.Printf("Кратчайший путь от города %s:\n", start)
	for city, dist := range distances {
		if dist == math.MaxInt {
			fmt.Printf("До города %s нет пути\n", city)
		} else {
			fmt.Printf("Расстояние до города %s: %d км\n", city, dist)
		}
	}
}

func main() {
	// Создаем граф
	g := NewGraph()

	// Добавляем дороги между городами (граф)
	g.AddEdge("A", "B", 10)
	g.AddEdge("A", "C", 3)
	g.AddEdge("B", "C", 1)
	g.AddEdge("B", "D", 2)
	g.AddEdge("C", "B", 4)
	g.AddEdge("C", "D", 8)
	g.AddEdge("C", "E", 2)
	g.AddEdge("D", "E", 7)
	g.AddEdge("E", "D", 9)

	// Выводим кратчайшие пути от города A
	g.PrintShortestPaths("E")
}
