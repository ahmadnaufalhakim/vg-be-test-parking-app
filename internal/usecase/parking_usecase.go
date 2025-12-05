package usecase

import (
	"container/heap"
	"fmt"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/domain"
)

func CreateParkingLot(capacity int) (*domain.ParkingLot, error) {
	return domain.NewParkingLot(capacity)
}

func ParkCar(
	parkingLot *domain.ParkingLot,
	carNumber string,
) (*domain.ParkingSlot, error) {
	return parkingLot.ParkCar(carNumber)
}

func Status(parkingLot *domain.ParkingLot) error {
	tmpHeap := &domain.ParkingHeap{}
	heap.Init(tmpHeap)

	for _, parkingSlot := range parkingLot.OccupiedSlots {
		heap.Push(tmpHeap, parkingSlot)
	}

	fmt.Printf("Slot No.\tRegistration No.\n")
	for range tmpHeap.Len() {
		parkingSlot, ok := heap.Pop(tmpHeap).(*domain.ParkingSlot)
		if !ok {
			return fmt.Errorf("ParkingSlot type assertion failed")
		}

		fmt.Printf("%d\t%s\n", parkingSlot.ID, parkingSlot.CarNumber)
	}

	return nil
}

func RemoveCar(
	parkingLot *domain.ParkingLot,
	carNumber string,
) (*domain.ParkingSlot, error) {
	return parkingLot.RemoveCar(carNumber)
}

func CalculateParkingCharge(hours int) (int, error) {
	if hours < 0 {
		return 0, fmt.Errorf("invalid hours input")
	}

	if hours < 2 {
		return 10, nil
	}

	return 10 * (hours - 1), nil
}
