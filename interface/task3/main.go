package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrUnsupported = errors.New("обновление недоступно")
)

type Device interface {
	UpdateOS(version string) error
	GetInfo() string
}

type Smartphone struct {
	OSVersion string
	Model     string
}

func NewSmartphone(OSVersion, Model string) *Smartphone {

	return &Smartphone{OSVersion: OSVersion, Model: Model}
}

func (s *Smartphone) UpdateOS(version string) error {
	osV, err := strconv.ParseFloat(s.OSVersion, 64)
	if err != nil {
		return err
	}
	if osV >= 12.0 {
		return ErrUnsupported
	}
	s.OSVersion = version
	return nil
}

func (s *Smartphone) GetInfo() string {
	return "Модель: " + s.Model + ", ОС: " + s.OSVersion
}

type Laptop struct {
	OSVersion string
	Model     string
}

func NewLaptop(OSVersion, Model string) *Laptop {

	return &Laptop{OSVersion: OSVersion, Model: Model}
}

func (l *Laptop) UpdateOS(version string) error {
	if !strings.HasPrefix(version, "Windows") {
		return ErrUnsupported
	}
	l.OSVersion = version
	return nil
}

func (l *Laptop) GetInfo() string {
	return "Модель: " + l.Model + ", ОС: " + l.OSVersion
}

type Smartwatch struct {
	OSVersion string
	Model     string
}

func NewSmartwatch(OSVersion, Model string) *Smartwatch {

	return &Smartwatch{OSVersion: OSVersion, Model: Model}
}

func (s *Smartwatch) UpdateOS(version string) error {
	if len(version) < 5 {
		return ErrUnsupported
	}
	s.OSVersion = version
	return nil
}
func (s *Smartwatch) GetInfo() string {
	return "Модель: " + s.Model + ", ОС: " + s.OSVersion
}

func main() {
	var laptop Device
	var Iphone Device
	var watch Device

	Iphone= NewSmartphone("12.0", "Iphone 17 Pro")
	laptop = NewLaptop("Windows 8", "MSI XZ17 PRO 24")
	watch = NewSmartwatch("OS511", "Generic Smartwatch 2")


	fmt.Println(laptop.GetInfo())
	fmt.Println(laptop.UpdateOS("9"))
	fmt.Println(laptop.GetInfo())
	fmt.Println(laptop.UpdateOS("Windows 10"))
	fmt.Println(laptop.GetInfo())

	fmt.Println(Iphone.GetInfo())
	fmt.Println(Iphone.UpdateOS("13.0"))

	fmt.Println(watch.GetInfo())
	fmt.Println(watch.UpdateOS("OSs"))
	fmt.Println(watch.GetInfo())
	fmt.Println(watch.UpdateOS("OS522"))
	fmt.Println(watch.GetInfo())

}
