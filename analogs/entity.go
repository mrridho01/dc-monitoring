package analogs

import "gorm.io/gorm"

type InputRegister struct {
	AcInput, Dc110, Dc48, CurrentAc, CurrentDc float32 `gorm:"default:0"`
}

type Analog struct {
	gorm.Model
	InputRegister InputRegister `gorm:"embedded"`
	GHID          uint
}
