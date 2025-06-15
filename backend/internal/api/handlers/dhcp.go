package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/freedyc/devin-app/backend/internal/models"
	"github.com/gin-gonic/gin"
)

var dhcpScopes = []models.DHCPScope{
	{
		ID:         1,
		Name:       "Main Network",
		Network:    "192.168.1.0",
		Netmask:    "255.255.255.0",
		RangeStart: "192.168.1.100",
		RangeEnd:   "192.168.1.200",
		Gateway:    "192.168.1.1",
		DNSServers: "8.8.8.8,8.8.4.4",
		LeaseTime:  86400,
		Enabled:    true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
}

var dhcpReservations = []models.DHCPReservation{
	{
		ID:          1,
		ScopeID:     1,
		MACAddress:  "00:11:22:33:44:55",
		IPAddress:   "192.168.1.50",
		Hostname:    "server1",
		Description: "Main server",
		Enabled:     true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

var dhcpLeases = []models.DHCPLease{
	{
		ID:         1,
		IPAddress:  "192.168.1.101",
		MACAddress: "aa:bb:cc:dd:ee:ff",
		Hostname:   "client1",
		StartTime:  time.Now().Add(-time.Hour),
		EndTime:    time.Now().Add(23 * time.Hour),
		State:      "active",
		ScopeID:    1,
	},
}

func GetDHCPScopes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"scopes": dhcpScopes,
		"total":  len(dhcpScopes),
	})
}

func CreateDHCPScope(c *gin.Context) {
	var req models.CreateScopeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scope := models.DHCPScope{
		ID:         len(dhcpScopes) + 1,
		Name:       req.Name,
		Network:    req.Network,
		Netmask:    req.Netmask,
		RangeStart: req.RangeStart,
		RangeEnd:   req.RangeEnd,
		Gateway:    req.Gateway,
		DNSServers: req.DNSServers,
		LeaseTime:  req.LeaseTime,
		Enabled:    true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if scope.LeaseTime == 0 {
		scope.LeaseTime = 86400
	}

	dhcpScopes = append(dhcpScopes, scope)
	c.JSON(http.StatusCreated, scope)
}

func GetDHCPScope(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scope ID"})
		return
	}

	for _, scope := range dhcpScopes {
		if scope.ID == id {
			c.JSON(http.StatusOK, scope)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Scope not found"})
}

func UpdateDHCPScope(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scope ID"})
		return
	}

	var req models.CreateScopeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, scope := range dhcpScopes {
		if scope.ID == id {
			dhcpScopes[i].Name = req.Name
			dhcpScopes[i].Network = req.Network
			dhcpScopes[i].Netmask = req.Netmask
			dhcpScopes[i].RangeStart = req.RangeStart
			dhcpScopes[i].RangeEnd = req.RangeEnd
			dhcpScopes[i].Gateway = req.Gateway
			dhcpScopes[i].DNSServers = req.DNSServers
			dhcpScopes[i].LeaseTime = req.LeaseTime
			dhcpScopes[i].UpdatedAt = time.Now()
			c.JSON(http.StatusOK, dhcpScopes[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Scope not found"})
}

func DeleteDHCPScope(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scope ID"})
		return
	}

	for i, scope := range dhcpScopes {
		if scope.ID == id {
			dhcpScopes = append(dhcpScopes[:i], dhcpScopes[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Scope deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Scope not found"})
}

func GetDHCPReservations(c *gin.Context) {
	scopeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scope ID"})
		return
	}

	var reservations []models.DHCPReservation
	for _, reservation := range dhcpReservations {
		if reservation.ScopeID == scopeID {
			reservations = append(reservations, reservation)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"reservations": reservations,
		"total":        len(reservations),
	})
}

func CreateDHCPReservation(c *gin.Context) {
	scopeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scope ID"})
		return
	}

	var req models.CreateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservation := models.DHCPReservation{
		ID:          len(dhcpReservations) + 1,
		ScopeID:     scopeID,
		MACAddress:  req.MACAddress,
		IPAddress:   req.IPAddress,
		Hostname:    req.Hostname,
		Description: req.Description,
		Enabled:     true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	dhcpReservations = append(dhcpReservations, reservation)
	c.JSON(http.StatusCreated, reservation)
}

func GetDHCPReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}

	for _, reservation := range dhcpReservations {
		if reservation.ID == id {
			c.JSON(http.StatusOK, reservation)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
}

func UpdateDHCPReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}

	var req models.CreateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, reservation := range dhcpReservations {
		if reservation.ID == id {
			dhcpReservations[i].MACAddress = req.MACAddress
			dhcpReservations[i].IPAddress = req.IPAddress
			dhcpReservations[i].Hostname = req.Hostname
			dhcpReservations[i].Description = req.Description
			dhcpReservations[i].UpdatedAt = time.Now()
			c.JSON(http.StatusOK, dhcpReservations[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
}

func DeleteDHCPReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}

	for i, reservation := range dhcpReservations {
		if reservation.ID == id {
			dhcpReservations = append(dhcpReservations[:i], dhcpReservations[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
}

func GetDHCPLeases(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"leases": dhcpLeases,
		"total":  len(dhcpLeases),
	})
}

func GetDHCPLeasesByScope(c *gin.Context) {
	scopeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scope ID"})
		return
	}

	var leases []models.DHCPLease
	for _, lease := range dhcpLeases {
		if lease.ScopeID == scopeID {
			leases = append(leases, lease)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"leases": leases,
		"total":  len(leases),
	})
}
