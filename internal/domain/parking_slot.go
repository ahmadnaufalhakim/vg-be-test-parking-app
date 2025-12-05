package domain

type ParkingSlot struct {
	ID                int
	DistanceFromEntry int
	Available         bool
	CarNumber         string
}
