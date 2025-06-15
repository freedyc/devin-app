package models

import (
	"time"
)

type GSLBPolicy struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	FQDN        string    `json:"fqdn" db:"fqdn"`
	Method      string    `json:"method" db:"method"`
	TTL         int       `json:"ttl" db:"ttl"`
	Enabled     bool      `json:"enabled" db:"enabled"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type GSLBPool struct {
	ID          int       `json:"id" db:"id"`
	PolicyID    int       `json:"policy_id" db:"policy_id"`
	Name        string    `json:"name" db:"name"`
	Priority    int       `json:"priority" db:"priority"`
	Weight      int       `json:"weight" db:"weight"`
	Enabled     bool      `json:"enabled" db:"enabled"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type GSLBServer struct {
	ID          int       `json:"id" db:"id"`
	PoolID      int       `json:"pool_id" db:"pool_id"`
	Name        string    `json:"name" db:"name"`
	IPAddress   string    `json:"ip_address" db:"ip_address"`
	Port        int       `json:"port" db:"port"`
	Weight      int       `json:"weight" db:"weight"`
	Priority    int       `json:"priority" db:"priority"`
	HealthCheck string    `json:"health_check" db:"health_check"`
	Status      string    `json:"status" db:"status"`
	Enabled     bool      `json:"enabled" db:"enabled"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreatePolicyRequest struct {
	Name        string `json:"name" binding:"required"`
	FQDN        string `json:"fqdn" binding:"required"`
	Method      string `json:"method" binding:"required"`
	TTL         int    `json:"ttl"`
	Description string `json:"description"`
}

type CreatePoolRequest struct {
	PolicyID int    `json:"policy_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Priority int    `json:"priority"`
	Weight   int    `json:"weight"`
}

type CreateServerRequest struct {
	PoolID      int    `json:"pool_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	IPAddress   string `json:"ip_address" binding:"required"`
	Port        int    `json:"port"`
	Weight      int    `json:"weight"`
	Priority    int    `json:"priority"`
	HealthCheck string `json:"health_check"`
}
