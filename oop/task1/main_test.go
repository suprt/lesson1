package main

import (
	"errors"
	"testing"
)

func TestVehiclePolymorphism(t *testing.T) {
	vehicles := []Vehicle{
		&Car{"Volvo", false},
		&Truck{Car{"KAMAZ", false}, 15.0},
		&ElectricCar{Car{"Haval", false}, 100},
	}

	for _, vehicle := range vehicles {
		if err := vehicle.StartEngine(); err != nil {
			t.Errorf("StartEngine() error = %v", err)
		}
		if err := vehicle.StopEngine(); err != nil {
			t.Errorf("StopEngine() error = %v", err)
		}
	}
}

func TestStartEngine(t *testing.T) {
	tests := []struct {
		name    string
		input   Vehicle
		wantErr error
	}{
		{
			name:    "CarSuccessStart",
			input:   &Car{"Volvo", false},
			wantErr: nil,
		},
		{
			name:    "CarErrorStart",
			input:   &Car{"Volvo", true},
			wantErr: ErrEngineAlreadyRunning,
		},
		{
			name:    "TruckSuccessStart",
			input:   &Truck{Car{"KAMAZ", false}, 15.0},
			wantErr: nil,
		},
		{
			name:    "TruckErrorStart",
			input:   &Truck{Car{"KAMAZ", true}, 15.0},
			wantErr: ErrEngineAlreadyRunning,
		},
		{
			name:    "ECarSuccessStart",
			input:   &ElectricCar{Car{"Haval", false}, 100},
			wantErr: nil,
		},
		{
			name:    "ECarErrorEngineStart",
			input:   &ElectricCar{Car{"Haval", true}, 100},
			wantErr: ErrEngineAlreadyRunning,
		},
		{
			name:    "ECarErrorBatteryStart",
			input:   &ElectricCar{Car{"Haval", true}, 5},
			wantErr: ErrLowBattery,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := tt.input.StartEngine(); !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStopEngine(t *testing.T) {
	tests := []struct {
		name    string
		input   Vehicle
		wantErr error
	}{
		{
			name:    "CarSuccessStop",
			input:   &Car{"Volvo", true},
			wantErr: nil,
		},
		{
			name:    "CarErrorStop",
			input:   &Car{"Volvo", false},
			wantErr: ErrEngineOff,
		},
		{
			name:    "TruckSuccessStop",
			input:   &Truck{Car{"KAMAZ", true}, 15.0},
			wantErr: nil,
		},
		{
			name:    "TruckErrorStop",
			input:   &Truck{Car{"KAMAZ", false}, 15.0},
			wantErr: ErrEngineOff,
		},
		{
			name:    "ECarSuccessStop",
			input:   &ElectricCar{Car{"Haval", true}, 100},
			wantErr: nil,
		},
		{
			name:    "ECarErrorEngineStop",
			input:   &ElectricCar{Car{"Haval", false}, 100},
			wantErr: ErrEngineOff,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := tt.input.StopEngine(); !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHonk(t *testing.T) {
	tests := []struct {
		name  string
		input interface {
			Honk() string
		}
		want string
	}{
		{
			name:  "CarHonk",
			input: &Car{},
			want:  "Beep beep!",
		},
		{
			name:  "TruckHonk",
			input: &Truck{},
			want:  "Honk Honk!",
		},
		{
			name:  "ECarHonk",
			input: &ElectricCar{},
			want:  "Beep beep!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Honk()
			if got != tt.want {
				t.Errorf("Honk() = %q, want %q", got, tt.want)
			}
		})
	}
}
