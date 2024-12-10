package main

import "fmt"

// Интерфейс нового стандарта (USB-C)
type USBC interface {
	ConnectUSBC()
}

// Старое устройство с интерфейсом USB-A
type USBDevice struct{}

func (d *USBDevice) ConnectUSBA() {
	fmt.Println("Подключено через USB-A")
}

// Адаптер для USB-A -> USB-C
type USBAdapter struct {
	device *USBDevice
}

// Реализация нового интерфейса через адаптер
func (a *USBAdapter) ConnectUSBC() {
	fmt.Println("Адаптер преобразует сигнал USB-A в USB-C")
	a.device.ConnectUSBA()
}

// Новая система, работающая только с USB-C
func charge(device USBC) {
	fmt.Println("Зарядка через USB-C началась")
	device.ConnectUSBC()
}

func main() {
	// Старое устройство с USB-A
	oldDevice := &USBDevice{}

	// Подключаем адаптер
	adapter := &USBAdapter{device: oldDevice}

	// Новая система использует адаптер для зарядки старого устройства
	charge(adapter)
}
