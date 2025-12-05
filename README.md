# vg-be-test-parking-app
CLI-based parking app (automated ticketing system) written in Go. The program reads commands from a file and executes simple parking-lot operations (create, park, leave, status).

## Quick start

Build:
```sh
go build ./cmd/parking-app
```

Run with an input file:
```sh
./parking-app input/example-1.txt
# or
go run ./cmd/parking-app/main.go input/example-1.txt
```

See sample commands: [input/example-1.txt](input/example-1.txt)

## Commands
- `create_parking_lot <capacity>`
- `park <car_number>`
- `leave <car_number> <hours>`
- `status`

Command parsing is performed by [`io.ExtractCommandsFromFile`](internal/io/file_parser.go) and dispatched via [`cli.NewDispatcher`](internal/cli/dispatcher.go) / [`cli.Dispatcher.Handle`](internal/cli/dispatcher.go).

## Key files & symbols
- Entry point: [cmd/parking-app/main.go](cmd/parking-app/main.go) — uses [`io.ExtractCommandsFromFile`](internal/io/file_parser.go) and dispatcher.
- CLI router and handlers: [internal/cli/dispatcher.go](internal/cli/dispatcher.go), [internal/cli/handlers.go](internal/cli/handlers.go) (see `createParkingLot`, `park`, `leave`, `status`).
  - [`cli.NewDispatcher`](internal/cli/dispatcher.go)
- Domain model and heap:
  - [`domain.NewParkingLot`](internal/domain/parking_lot.go)
  - [`domain.ParkingLot`](internal/domain/parking_lot.go)
  - [`domain.ParkingSlot`](internal/domain/parking_slot.go)
  - [`domain.ParkingHeap`](internal/domain/heap.go)
- Use-cases / business logic: [internal/usecase/parking_usecase.go](internal/usecase/parking_usecase.go)
  - [`usecase.CreateParkingLot`](internal/usecase/parking_usecase.go)
  - [`usecase.ParkCar`](internal/usecase/parking_usecase.go)
  - [`usecase.RemoveCar`](internal/usecase/parking_usecase.go)
  - [`usecase.CalculateParkingCharge`](internal/usecase/parking_usecase.go)
  - [`usecase.Status`](internal/usecase/parking_usecase.go)
- Utilities: error formatting in [internal/utils/error.go](internal/utils/error.go)

## Notes
- Priority queue uses Go's [`container/heap`](https://pkg.go.dev/container/heap) implementation (see [internal/domain/heap.go](internal/domain/heap.go)).
- The app prints user-facing messages from handlers in [internal/cli/handlers.go](internal/cli/handlers.go).
- Input file must contain one command per line (see [input/example-1.txt](input/example-1.txt)).

License: MIT — see [LICENSE](LICENSE)