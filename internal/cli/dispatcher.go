package cli

import (
	"fmt"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/domain"
)

type Dispatcher struct {
	parkingLot *domain.ParkingLot
	handlers   map[string]func([]string) error
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		parkingLot: nil,
		handlers:   make(map[string]func([]string) error),
	}

	// Router-like mapping from command to handler
	d.handlers["create_parking_lot"] = d.createParkingLot
	d.handlers["park"] = d.park
	d.handlers["leave"] = d.leave
	d.handlers["status"] = d.status

	return d
}

func (d *Dispatcher) Handle(cmd string, args []string) error {
	handler, ok := d.handlers[cmd]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd)
	}

	return handler(args)
}

// The following below are handler function wrappers for testing

func (d *Dispatcher) CreateParkingLot(args []string) error { return d.createParkingLot(args) }
func (d *Dispatcher) Park(args []string) error             { return d.park(args) }
func (d *Dispatcher) Leave(args []string) error            { return d.leave(args) }
func (d *Dispatcher) Status(args []string) error           { return d.status(args) }

// Public parkingLot getter for testing
func (d *Dispatcher) GetParkingLot() *domain.ParkingLot { return d.parkingLot }
