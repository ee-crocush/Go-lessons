package main

import (
	"fmt"
	// "strconv"
)

// Представляет узел односвязного списка (соседа вершины)
type Node struct {
	vertex int   // Вершина
	next   *Node // Сосед
}

// Представляет вершину графа
type Vertex struct {
	number    int   // Номер вершины
	neighbors *Node // Список соседей
}

// Представляет неориентированный граф
type Graph struct {
	vertices []*Vertex // Список вершин
}

// Конструктор для создания нового графа
func NewGraph() *Graph {
	return &Graph{
		vertices: []*Vertex{},
	}
}

// Добавляет вершину в граф
func (g *Graph) AddVertex(number int) {
	for _, v := range g.vertices {
		if v.number == number {
			fmt.Printf("Вершина %d уже существует\n", number)
			return
		}
	}
	g.vertices = append(g.vertices, &Vertex{number: number})
}

// Добавляет ребро между двумя вершинами
func (g *Graph) AddEdge(v1, v2 int) {
	v1Vertex := g.getVertex(v1)
	v2Vertex := g.getVertex(v2)

	if v1Vertex == nil || v2Vertex == nil {
		fmt.Println("Одна или обе вершины не существуют")
		return
	}

	// Добавляем ребро v1 -> v2
	v1Vertex.neighbors = &Node{vertex: v2, next: v1Vertex.neighbors}

	// Добавляем ребро v2 -> v1 (неориентированный граф)
	v2Vertex.neighbors = &Node{vertex: v1, next: v2Vertex.neighbors}
}

// Находит вершину по её номеру
func (g *Graph) getVertex(number int) *Vertex {
	for _, v := range g.vertices {
		if v.number == number {
			return v
		}
	}
	return nil
}

// Выводит граф
func (g *Graph) PrintGraph() {
	for _, v := range g.vertices {
		fmt.Printf("%d ->", v.number)
		neighbor := v.neighbors
		for neighbor != nil {
			fmt.Printf(" %d", neighbor.vertex)
			neighbor = neighbor.next
		}
		fmt.Println()
	}
}

// Выполняет поиск в ширину и возвращает маршрут
func (g *Graph) BFSWithPath(start, target int) {
	startVertex := g.getVertex(start)

	if startVertex == nil {
		fmt.Printf("Вершина %d не найдена\n", start)
		return
	}

	visited := make(map[int]bool)   // Множество посещённых вершин
	queue := []*Vertex{startVertex} // Очередь посещений
	previouses := make(map[int]int) // Карта предшественников

	found := false // Флаг наличия пути

	// BFS
	for len(queue) > 0 {
		// Удаляем первую вершину из очереди
		current := queue[0]
		queue = queue[1:]

		// Если вершина уже была посещена, пропускаем её
		if visited[current.number] {
			continue
		}

		// Помечаем вершину как посещённую
		visited[current.number] = true

		// Если нашли целевую вершину
		if current.number == target {
			found = true
			break
		}

		// Добавляем всех соседей текущей вершины в очередь
		neighbor := current.neighbors
		for neighbor != nil {
			if !visited[neighbor.vertex] {
				queue = append(queue, g.getVertex(neighbor.vertex))
				// Записываем предшественника
				if _, exists := previouses[neighbor.vertex]; !exists {
					previouses[neighbor.vertex] = current.number
				}
			}
			neighbor = neighbor.next
		}
	}

	if !found {
		fmt.Printf("Целевая вершина %d недостижима из вершины %d\n", target, start)
		return
	}

	// Восстанавливаем маршрут
	path := g.reconstructPath(previouses, start, target)

	PrintPath(path, start, target)
}

// Восстанавливает маршрут от начальной вершины до целевой
func (g *Graph) reconstructPath(previouses map[int]int, start, target int) []int {
	path := []int{}
	for at := target; at != start; at = previouses[at] {
		path = append([]int{at}, path...)
	}
	path = append([]int{start}, path...) // Добавляем стартовую вершину
	return path
}

func PrintPath(path []int, start, target int) {
	var prettyPath string
	for i, p := range path {
		// Для последней вершины не добавляем стрелку
		if i == len(path)-1 {
			prettyPath += fmt.Sprintf("%d", p)
		} else {
			prettyPath += fmt.Sprintf("%d->", p)
		}
	}
	fmt.Printf("Маршрут от %d до %d: %v\n", start, target, prettyPath)
}

func main() {
	// Создаем граф через конструктор
	graph := NewGraph()

	// Добавляем вершины
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)

	// Добавляем рёбра
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 5)
	graph.AddEdge(4, 5)

	// Выводим граф
	graph.PrintGraph()

	// Выполняем поиск в ширину с восстановлением пути
	graph.BFSWithPath(1, 5)
	graph.BFSWithPath(1, 6) // Проверка для недостижимой вершины
}
