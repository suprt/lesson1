package main

import (
	"errors"
	"fmt"
)

var (
	ErrEngineAlreadyRunning = errors.New("двигатель уже работает")
	ErrEngineOff            = errors.New("двигатель не запущен")
	ErrLowBattery           = errors.New("низкий заряд батареи")
)

type Vehicle interface {
	StartEngine() error
	StopEngine() error
	GetInfo() string
}

type Car struct {
	Brand    string
	engineOn bool
}

func (c *Car) Honk() string {
	return "Beep beep!"
}
func (c *Car) StartEngine() error {
	if c.engineOn {
		return ErrEngineAlreadyRunning
	}
	c.engineOn = true
	return nil
}
func (c *Car) StopEngine() error {
	if !c.engineOn {
		return ErrEngineOff
	}
	c.engineOn = false
	return nil
}
func (c *Car) GetInfo() string {
	return fmt.Sprintf("Brand: %s, EnginOn: %t", c.Brand, c.engineOn)
}
func (c *Car) GetEngineStatus() bool {
	return c.engineOn
}

type Truck struct {
	Car
	cargoCapacity float64
}

func (t *Truck) Honk() string {
	return "Honk Honk!"
}
func (t *Truck) GetInfo() string {
	return fmt.Sprintf("Brand: %s, EnginOn: %t, Cargo capacity: %g", t.Brand, t.engineOn, t.cargoCapacity)
}
func (t *Truck) GetCargoCapacity() float64 {
	return t.cargoCapacity
}

type ElectricCar struct {
	Car
	batteryLevel int
}

func (e *ElectricCar) StartEngine() error {
	if e.batteryLevel <= 5 {
		return ErrLowBattery
	}

	return e.Car.StartEngine()
}
func (e *ElectricCar) GetInfo() string {
	return fmt.Sprintf("Brand: %s, EnginOn: %t, Battery level: %v%%", e.Brand, e.engineOn, e.batteryLevel)
}
func (e *ElectricCar) GetBatteryLevel() int {
	return e.batteryLevel
}

func main() {

}
