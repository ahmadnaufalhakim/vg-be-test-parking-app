package domain

// ParkingHeap implements heap interface from
// Golang's built-in heap library
type ParkingHeap []*ParkingSlot

func (h ParkingHeap) Len() int {
	return len(h)
}

func (h ParkingHeap) Less(i, j int) bool {
	return h[i].DistanceFromEntry < h[j].DistanceFromEntry
}

func (h ParkingHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ParkingHeap) Push(x any) {
	*h = append(*h, x.(*ParkingSlot))
}

func (h *ParkingHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
