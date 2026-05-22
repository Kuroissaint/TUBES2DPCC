package main
type DispatchRepository interface {
    UpdateDriverStatus(driverID string, status string) error
}