package graph

// Matrix - представление графа в виде матрицы смежности.
// Matrix[i][j] содержит стоимость ребра между вершинами i и j.
type Matrix [][]int

// List - представление графа в виде списка смежности.
// Лист смежности содержит для каждой вершины графа лист
// смежных вершин, с указанием стоимости ребра.
type List map[int][]Neighbour

// Neighbour - смежная вершина и вес ребра.
type Neighbour struct {
	Vertex int
	Weight int
}

// NewList создаёт новый список смежности.
func NewList() List {
	var l List
	l = make(map[int][]Neighbour)
	return l
}

// AddNodes добавляет вершины графа.
func (l List) AddNodes(nodes ...int) {
	for _, node := range nodes {
		l[node] = []Neighbour{}
	}
}

// AddEdge добавляет ребро графа.
func (l List) AddEdge(node1, node2, weight int) {
	l[node1] = append(l[node1], Neighbour{Vertex: node2, Weight: weight})
}

// Weight возвращает стоимость всех рёбер графа.
func (l List) Weight() int {
	sum := 0
	for _, v := range l {
		for _, edge := range v {
			sum += edge.Weight
		}
	}
	return sum
}

// MinSpanTree - поиск минимального покрывающего (остовного) дерева
// алгоритмом Прима.
func (l List) MinSpanTree() List {

	// алгоритм:
	// - определяем первую вершину и добавляем в исследуемые
	// - цикл для исследуемых вершин:
	// 		- находим ближайшего соседа и добавляем ребро в дерево

	tree := NewList()
	current := make(map[int]bool) // вершины, исследуемые на данном шаге
	current[1] = true
	visited := make(map[int]bool) // посещённые вершины

	initVal := func() (v1, v2, dist int) {
		for k, v := range l {
			if k > v1 {
				v1, v2 = k, k
			}
			for _, item := range v {
				if item.Weight > dist {
					dist = item.Weight
				}
			}
		}
		return v1 + 1, v2 + 1, dist + 1
	}
	initV, _, _ := initVal()

OUT:
	for {
		v1, v2, dist := initVal()

		for i := range current {
			v := l[i]
			if len(v) == 0 {
				continue
			}
			for _, edge := range v {
				if visited[edge.Vertex] {
					continue
				}

				if edge.Weight < dist {
					v1 = i
					v2 = edge.Vertex
					dist = edge.Weight
				}
			}
		}
		if v1 == initV {
			break OUT
		}
		visited[v1] = true
		visited[v2] = true
		neighbour := Neighbour{Vertex: v2, Weight: dist}
		if _, ok := tree[v1]; !ok {
			tree[v1] = []Neighbour{}
		}
		tree[v1] = append(tree[v1], neighbour)
		edges := 0
		for _, v := range tree {
			edges += len(v)
		}
		if edges+1 == len(l) {
			break OUT
		}
		current[v2] = true
	}
	return tree
}
