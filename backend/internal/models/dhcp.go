package models

import (
	"time"
)

type DHCPScope struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Network     string    `json:"network" db:"network"`
	Netmask     string    `json:"netmask" db:"netmask"`
	RangeStart  string    `json:"range_start" db:"range_start"`
	RangeEnd    string    `json:"range_end" db:"range_end"`
	Gateway     string    `json:"gateway" db:"gateway"`
	DNSServers  string    `json:"dns_servers" db:"dns_servers"`
	LeaseTime   int       `json:"lease_time" db:"lease_time"`
	Enabled     bool      `json:"enabled" db:"enabled"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type DHCPReservation struct {
	ID          int       `json:"id" db:"id"`
	ScopeID     int       `json:"scope_id" db:"scope_id"`
	MACAddress  string    `json:"mac_address" db:"mac_address"`
	IPAddress   string    `json:"ip_address" db:"ip_address"`
	Hostname    string    `json:"hostname" db:"hostname"`
	Description string    `json:"description" db:"description"`
	Enabled     bool      `json:"enabled" db:"enabled"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type DHCPLease struct {
	ID         int       `json:"id" db:"id"`
	IPAddress  string    `json:"ip_address" db:"ip_address"`
	MACAddress string    `json:"mac_address" db:"mac_address"`
	Hostname   string    `json:"hostname" db:"hostname"`
	StartTime  time.Time `json:"start_time" db:"start_time"`
	EndTime    time.Time `json:"end_time" db:"end_time"`
	State      string    `json:"state" db:"state"`
	ScopeID    int       `json:"scope_id" db:"scope_id"`
}

type CreateScopeRequest struct {
	Name       string `json:"name" binding:"required"`
	Network    string `json:"network" binding:"required"`
	Netmask    string `json:"netmask" binding:"required"`
	RangeStart string `json:"range_start" binding:"required"`
	RangeEnd   string `json:"range_end" binding:"required"`
	Gateway    string `json:"gateway" binding:"required"`
	DNSServers string `json:"dns_servers"`
	LeaseTime  int    `json:"lease_time"`
}

type CreateReservationRequest struct {
	ScopeID     int    `json:"scope_id" binding:"required"`
	MACAddress  string `json:"mac_address" binding:"required"`
	IPAddress   string `json:"ip_address" binding:"required"`
	Hostname    string `json:"hostname"`
	Description string `json:"description"`
}
