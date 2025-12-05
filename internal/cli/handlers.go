package cli

import (
	"fmt"
	"strconv"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/usecase"
)

func (d *Dispatcher) createParkingLot(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: create_parking_lot <capacity>")
	}

	capacity, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	parkingLot, err := usecase.CreateParkingLot(capacity)
	if err != nil {
		return err
	}

	d.parkingLot = parkingLot
	return nil
}

func (d *Dispatcher) park(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: park <car_number>")
	}

	if d.parkingLot == nil {
		return fmt.Errorf("parking lot not initialized")
	}

	parkingSlot, err := usecase.ParkCar(d.parkingLot, args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Allocated slot number: %d\n", parkingSlot.ID)
	return nil
}

func (d *Dispatcher) leave(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: leave <car_number> <hours>")
	}

	if d.parkingLot == nil {
		return fmt.Errorf("parking lot not initialized")
	}

	hours, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	parkingSlot, err := usecase.RemoveCar(d.parkingLot, args[0])
	if err != nil {
		return err
	}

	parkingCharge, err := usecase.CalculateParkingCharge(hours)
	if err != nil {
		return err
	}

	fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n",
		args[0], parkingSlot.ID, parkingCharge)
	return nil
}

func (d *Dispatcher) status(args []string) error {
	if d.parkingLot == nil {
		return fmt.Errorf("parking lot not initialized")
	}

	return usecase.Status(d.parkingLot)
}
