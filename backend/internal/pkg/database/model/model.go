package model

import (
	"time"
)

type Registrant struct {
	VerificationCode string `json:"code" xml:"code" form:"code" redis:"code"`
	Name             string `json:"name" xml:"name" form:"name" redis:"name"`
	PhoneNumber      string `json:"phonenumber" xml:"phonenumber" form:"phonenumber" redis:"phonenumber"`
	Password         string `json:"password" xml:"password" form:"password" redis:"password"`
}

type User struct {
	ID              uint              `json:"id" xml:"id" form:"id" gorm:"primarykey;index"`
	CreatedAt       time.Time         `json:"createdAt" xml:"createdAt" form:"createdAt"`
	UpdatedAt       time.Time         `json:"updatedAt" xml:"updatedAt" form:"updatedAt"`
	Name            string            `json:"name" xml:"name" form:"name"`
	PhoneNumber     string            `json:"phonenumber" xml:"phonenumber" form:"phonenumber"`
	Password        string            `json:"-" xml:"-" form:"-"`
	Farmlands       []Farmland        `json:"-" xml:"-" form:"-"`
	Microcontroller []Microcontroller `json:"-" xml:"-" form:"-"`
}

type UserSession struct {
	ID          uint   `json:"id" xml:"id" form:"id" redis:"id"`
	Name        string `json:"name" xml:"name" form:"name" redis:"name"`
	PhoneNumber string `json:"phonenumber" xml:"phonenumber" form:"phonenumber" redis:"phonenumber"`
}

type UserResetPassword struct {
	VerificationCode string `json:"code" xml:"code" form:"code" redis:"code"`
	PhoneNumber      string `json:"phonenumber" xml:"phonenumber" form:"phonenumber" redis:"phonenumber"`
}

type Farmland struct {
	ID               uint              `json:"id" xml:"id" form:"id" gorm:"primarykey;index"`
	CreatedAt        time.Time         `json:"createdAt" xml:"createdAt" form:"createdAt"`
	UpdatedAt        time.Time         `json:"updatedAt" xml:"updatedAt" form:"updatedAt"`
	Name             string            `json:"name" xml:"name" form:"name"`
	Description      string            `json:"description" xml:"description" form:"description"`
	Location         string            `json:"location" xml:"location" form:"location"`
	Microcontrollers []Microcontroller `json:"-" xml:"-" form:"-"`
	UserID           uint              `json:"-" xml:"-" form:"-"`
}

type Microcontroller struct {
	ID          uint       `json:"id" xml:"id" form:"id" gorm:"primarykey;index"`
	CreatedAt   time.Time  `json:"createdAt" xml:"createdAt" form:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt" xml:"updatedAt" form:"updatedAt"`
	Name        string     `json:"name" xml:"name" form:"name"`
	Description string     `json:"description" xml:"description" form:"description"`
	Location    string     `json:"location" xml:"location" form:"location"`
	DeviceID    string     `json:"deviceId" xml:"deviceId" form:"deviceId"`
	Snapshots   []Snapshot `json:"-" xml:"-" form:"-"`
	FarmlandID  uint       `json:"-" xml:"-" form:"-"`
	UserID      uint       `json:"-" xml:"-" form:"-"`
}

type Snapshot struct {
	ID                uint      `json:"-" xml:"-" form:"-" gorm:"primarykey;index"`
	CreatedAt         time.Time `json:"createdAt" xml:"createdAt" form:"createdAt"`
	Type              string    `json:"-" xml:"-" form:"-"`
	Data              string    `json:"data" xml:"data" form:"data"`
	Durations         []string  `json:"-" xml:"-" form:"-" gorm:"type:text[]"`
	DeviceID          string    `json:"-" xml:"-" form:"-"`
	MicrocontrollerID uint      `json:"-" xml:"-" form:"-"`
}
