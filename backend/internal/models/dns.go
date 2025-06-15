package models

import (
	"time"
)

type DNSZone struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Type        string    `json:"type" db:"type"`
	Serial      int       `json:"serial" db:"serial"`
	Refresh     int       `json:"refresh" db:"refresh"`
	Retry       int       `json:"retry" db:"retry"`
	Expire      int       `json:"expire" db:"expire"`
	MinTTL      int       `json:"min_ttl" db:"min_ttl"`
	PrimaryNS   string    `json:"primary_ns" db:"primary_ns"`
	AdminEmail  string    `json:"admin_email" db:"admin_email"`
	Enabled     bool      `json:"enabled" db:"enabled"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}



type DNSRecord struct {
	ID       int    `json:"id" db:"id"`
	ZoneID   int    `json:"zone_id" db:"zone_id"`
	Name     string `json:"name" db:"name"`
	Type     string `json:"type" db:"type"`
	Value    string `json:"value" db:"value"`
	TTL      int    `json:"ttl" db:"ttl"`
	Priority int    `json:"priority,omitempty" db:"priority"`
	Weight   int    `json:"weight,omitempty" db:"weight"`
	Port     int    `json:"port,omitempty" db:"port"`
	Enabled  bool   `json:"enabled" db:"enabled"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type DNSView struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	MatchClients string   `json:"match_clients" db:"match_clients"`
	Recursion   bool      `json:"recursion" db:"recursion"`
	Enabled     bool      `json:"enabled" db:"enabled"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type DNSZoneEnhanced struct {
	DNSZone
	ViewID      *int   `json:"view_id,omitempty" db:"view_id"`
	ZoneType    string `json:"zone_type" db:"zone_type"`
	Masters     string `json:"masters,omitempty" db:"masters"`
	Forwarders  string `json:"forwarders,omitempty" db:"forwarders"`
}

type ForwardingZone struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Forwarders []string  `json:"forwarders" db:"forwarders"`
	Forward    string    `json:"forward" db:"forward"`
	Enabled    bool      `json:"enabled" db:"enabled"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type StubZone struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Masters   []string  `json:"masters" db:"masters"`
	Enabled   bool      `json:"enabled" db:"enabled"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type RecursionConfig struct {
	ID              int      `json:"id" db:"id"`
	Enabled         bool     `json:"enabled" db:"enabled"`
	AllowRecursion  []string `json:"allow_recursion" db:"allow_recursion"`
	Forwarders      []string `json:"forwarders" db:"forwarders"`
	ForwardFirst    bool     `json:"forward_first" db:"forward_first"`
	DNSSECValidation string  `json:"dnssec_validation" db:"dnssec_validation"`
}

type ClusterNode struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	IPAddress string    `json:"ip_address" db:"ip_address"`
	Role      string    `json:"role" db:"role"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateZoneRequest struct {
	Name       string `json:"name" binding:"required"`
	Type       string `json:"type" binding:"required"`
	PrimaryNS  string `json:"primary_ns" binding:"required"`
	AdminEmail string `json:"admin_email" binding:"required"`
	Refresh    int    `json:"refresh"`
	Retry      int    `json:"retry"`
	Expire     int    `json:"expire"`
	MinTTL     int    `json:"min_ttl"`
}

type CreateRecordRequest struct {
	ZoneID   int    `json:"zone_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Value    string `json:"value" binding:"required"`
	TTL      int    `json:"ttl"`
	Priority int    `json:"priority,omitempty"`
	Weight   int    `json:"weight,omitempty"`
	Port     int    `json:"port,omitempty"`
}

type CreateViewRequest struct {
	Name         string `json:"name" binding:"required"`
	MatchClients string `json:"match_clients" binding:"required"`
	Recursion    bool   `json:"recursion"`
}

type CreateForwardingZoneRequest struct {
	Name       string   `json:"name" binding:"required"`
	Forwarders []string `json:"forwarders" binding:"required"`
	Forward    string   `json:"forward" binding:"required"`
}

type CreateStubZoneRequest struct {
	Name    string   `json:"name" binding:"required"`
	Masters []string `json:"masters" binding:"required"`
}
