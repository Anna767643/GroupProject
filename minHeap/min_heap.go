package minHeap

//- Push
//- Pop
//- IsEmpty
//- heapifyUp
//- heapifyDown

type Heap struct {
	data []Item
}

type Item struct {
	Vertex   int
	Distance int
}

func New() *Heap {
	return &Heap{}
}

func (h *Heap) Push(v Item) {

	h.data = append(h.data, v)

	i := len(h.data) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent].Distance <= h.data[i].Distance {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *Heap) Pop() (Item, bool) {
	if h.IsEmpty() {
		return Item{}, false
	}

	min := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	if h.IsEmpty() {
		return min, true
	}

	i := 0

	for {
		left := 2*i + 1
		right := 2*i + 2

		smallest := i

		// Проверяем левого ребёнка
		if left < len(h.data) && h.data[left].Distance < h.data[smallest].Distance {
			smallest = left
		}

		// Проверяем правого ребёнка
		if right < len(h.data) && h.data[right].Distance < h.data[smallest].Distance {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]

		i = smallest
	}
	return min, true
}

func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}
