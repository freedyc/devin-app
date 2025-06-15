package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/freedyc/devin-app/backend/internal/models"
	"github.com/gin-gonic/gin"
)

var dnsZones = []models.DNSZone{
	{
		ID:         1,
		Name:       "example.com",
		Type:       "master",
		Serial:     2024010101,
		Refresh:    3600,
		Retry:      1800,
		Expire:     604800,
		MinTTL:     86400,
		PrimaryNS:  "ns1.example.com",
		AdminEmail: "admin@example.com",
		Enabled:    true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
}

var dnsRecords = []models.DNSRecord{
	{
		ID:        1,
		ZoneID:    1,
		Name:      "@",
		Type:      "A",
		Value:     "192.168.1.100",
		TTL:       3600,
		Enabled:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        2,
		ZoneID:    1,
		Name:      "www",
		Type:      "A",
		Value:     "192.168.1.100",
		TTL:       3600,
		Enabled:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func GetDNSZones(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"zones": dnsZones,
		"total": len(dnsZones),
	})
}

func CreateDNSZone(c *gin.Context) {
	var req models.CreateZoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zone := models.DNSZone{
		ID:         len(dnsZones) + 1,
		Name:       req.Name,
		Type:       req.Type,
		Serial:     int(time.Now().Unix()),
		Refresh:    req.Refresh,
		Retry:      req.Retry,
		Expire:     req.Expire,
		MinTTL:     req.MinTTL,
		PrimaryNS:  req.PrimaryNS,
		AdminEmail: req.AdminEmail,
		Enabled:    true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if zone.Refresh == 0 {
		zone.Refresh = 3600
	}
	if zone.Retry == 0 {
		zone.Retry = 1800
	}
	if zone.Expire == 0 {
		zone.Expire = 604800
	}
	if zone.MinTTL == 0 {
		zone.MinTTL = 86400
	}

	dnsZones = append(dnsZones, zone)
	c.JSON(http.StatusCreated, zone)
}

func GetDNSZone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid zone ID"})
		return
	}

	for _, zone := range dnsZones {
		if zone.ID == id {
			c.JSON(http.StatusOK, zone)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Zone not found"})
}

func UpdateDNSZone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid zone ID"})
		return
	}

	var req models.CreateZoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, zone := range dnsZones {
		if zone.ID == id {
			dnsZones[i].Name = req.Name
			dnsZones[i].Type = req.Type
			dnsZones[i].PrimaryNS = req.PrimaryNS
			dnsZones[i].AdminEmail = req.AdminEmail
			dnsZones[i].Refresh = req.Refresh
			dnsZones[i].Retry = req.Retry
			dnsZones[i].Expire = req.Expire
			dnsZones[i].MinTTL = req.MinTTL
			dnsZones[i].UpdatedAt = time.Now()
			c.JSON(http.StatusOK, dnsZones[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Zone not found"})
}

func DeleteDNSZone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid zone ID"})
		return
	}

	for i, zone := range dnsZones {
		if zone.ID == id {
			dnsZones = append(dnsZones[:i], dnsZones[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Zone deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Zone not found"})
}

func GetDNSRecords(c *gin.Context) {
	zoneID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid zone ID"})
		return
	}

	var records []models.DNSRecord
	for _, record := range dnsRecords {
		if record.ZoneID == zoneID {
			records = append(records, record)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"records": records,
		"total":   len(records),
	})
}

func CreateDNSRecord(c *gin.Context) {
	zoneID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid zone ID"})
		return
	}

	var req models.CreateRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record := models.DNSRecord{
		ID:        len(dnsRecords) + 1,
		ZoneID:    zoneID,
		Name:      req.Name,
		Type:      req.Type,
		Value:     req.Value,
		TTL:       req.TTL,
		Priority:  req.Priority,
		Weight:    req.Weight,
		Port:      req.Port,
		Enabled:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if record.TTL == 0 {
		record.TTL = 3600
	}

	dnsRecords = append(dnsRecords, record)
	c.JSON(http.StatusCreated, record)
}

func GetDNSRecord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	for _, record := range dnsRecords {
		if record.ID == id {
			c.JSON(http.StatusOK, record)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
}

func UpdateDNSRecord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	var req models.CreateRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, record := range dnsRecords {
		if record.ID == id {
			dnsRecords[i].Name = req.Name
			dnsRecords[i].Type = req.Type
			dnsRecords[i].Value = req.Value
			dnsRecords[i].TTL = req.TTL
			dnsRecords[i].Priority = req.Priority
			dnsRecords[i].Weight = req.Weight
			dnsRecords[i].Port = req.Port
			dnsRecords[i].UpdatedAt = time.Now()
			c.JSON(http.StatusOK, dnsRecords[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
}

func DeleteDNSRecord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	for i, record := range dnsRecords {
		if record.ID == id {
			dnsRecords = append(dnsRecords[:i], dnsRecords[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
}
