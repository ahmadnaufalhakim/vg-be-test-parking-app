package internal

import (
	"container/heap"
	"fmt"
)

type ParkingLot struct {
	availableSlots *ParkingHeap
	occupiedSlots  map[string]*ParkingSlot
	capacity       int
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("capacity must be positive")
	}

	parkingHeap := &ParkingHeap{}
	heap.Init(parkingHeap)

	for i := range capacity {
		parkingSlot := &ParkingSlot{
			ID:                i + 1,
			DistanceFromEntry: i + 1,
			Available:         true,
			CarNumber:         "",
		}
		heap.Push(parkingHeap, parkingSlot)
	}

	return &ParkingLot{
		availableSlots: parkingHeap,
		occupiedSlots:  map[string]*ParkingSlot{},
		capacity:       capacity,
	}, nil
}

func (pl *ParkingLot) IsEmpty() bool {
	return pl.availableSlots.Len() == pl.capacity
}

func (pl *ParkingLot) IsFull() bool {
	return pl.availableSlots.Len() == 0
}

func (pl *ParkingLot) ParkCar(carNumber string) (*ParkingSlot, error) {
	if pl.IsFull() {
		return nil, fmt.Errorf("sorry, parking lot is full")
	}

	if parkingSlot, ok := pl.occupiedSlots[carNumber]; ok {
		return nil, fmt.Errorf(
			"registration number %s is already parked in Slot Number %d",
			parkingSlot.CarNumber, parkingSlot.ID,
		)
	}

	parkingSlot, ok := heap.Pop(pl.availableSlots).(*ParkingSlot)
	if !ok {
		return nil, fmt.Errorf("ParkingSlot type assertion failed")
	}

	parkingSlot.Available = false
	parkingSlot.CarNumber = carNumber
	pl.occupiedSlots[carNumber] = parkingSlot

	return parkingSlot, nil
}

func (pl *ParkingLot) RemoveCar(carNumber string) (*ParkingSlot, error) {
	if pl.IsEmpty() {
		return nil, fmt.Errorf("sorry, parking lot is empty")
	}

	parkingSlot, ok := pl.occupiedSlots[carNumber]
	if !ok {
		return nil, fmt.Errorf("registration number %s not found", carNumber)
	}

	parkingSlot.Available = true
	parkingSlot.CarNumber = ""

	heap.Push(pl.availableSlots, parkingSlot)
	delete(pl.occupiedSlots, carNumber)

	return parkingSlot, nil
}
