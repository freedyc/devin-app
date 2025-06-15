package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/freedyc/devin-app/backend/internal/models"
	"github.com/gin-gonic/gin"
)

var gslbPolicies = []models.GSLBPolicy{
	{
		ID:          1,
		Name:        "Web Load Balancer",
		FQDN:        "web.example.com",
		Method:      "weighted",
		TTL:         300,
		Enabled:     true,
		Description: "Load balancer for web servers",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

var gslbPools = []models.GSLBPool{
	{
		ID:        1,
		PolicyID:  1,
		Name:      "Primary Pool",
		Priority:  1,
		Weight:    100,
		Enabled:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

var gslbServers = []models.GSLBServer{
	{
		ID:          1,
		PoolID:      1,
		Name:        "Web Server 1",
		IPAddress:   "192.168.1.10",
		Port:        80,
		Weight:      100,
		Priority:    1,
		HealthCheck: "http",
		Status:      "healthy",
		Enabled:     true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

func GetGSLBPolicies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"policies": gslbPolicies,
		"total":    len(gslbPolicies),
	})
}

func CreateGSLBPolicy(c *gin.Context) {
	var req models.CreatePolicyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	policy := models.GSLBPolicy{
		ID:          len(gslbPolicies) + 1,
		Name:        req.Name,
		FQDN:        req.FQDN,
		Method:      req.Method,
		TTL:         req.TTL,
		Description: req.Description,
		Enabled:     true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if policy.TTL == 0 {
		policy.TTL = 300
	}

	gslbPolicies = append(gslbPolicies, policy)
	c.JSON(http.StatusCreated, policy)
}

func GetGSLBPolicy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid policy ID"})
		return
	}

	for _, policy := range gslbPolicies {
		if policy.ID == id {
			c.JSON(http.StatusOK, policy)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Policy not found"})
}

func UpdateGSLBPolicy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid policy ID"})
		return
	}

	var req models.CreatePolicyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, policy := range gslbPolicies {
		if policy.ID == id {
			gslbPolicies[i].Name = req.Name
			gslbPolicies[i].FQDN = req.FQDN
			gslbPolicies[i].Method = req.Method
			gslbPolicies[i].TTL = req.TTL
			gslbPolicies[i].Description = req.Description
			gslbPolicies[i].UpdatedAt = time.Now()
			c.JSON(http.StatusOK, gslbPolicies[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Policy not found"})
}

func DeleteGSLBPolicy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid policy ID"})
		return
	}

	for i, policy := range gslbPolicies {
		if policy.ID == id {
			gslbPolicies = append(gslbPolicies[:i], gslbPolicies[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Policy deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Policy not found"})
}

func GetGSLBPools(c *gin.Context) {
	policyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid policy ID"})
		return
	}

	var pools []models.GSLBPool
	for _, pool := range gslbPools {
		if pool.PolicyID == policyID {
			pools = append(pools, pool)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"pools": pools,
		"total": len(pools),
	})
}

func CreateGSLBPool(c *gin.Context) {
	policyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid policy ID"})
		return
	}

	var req models.CreatePoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pool := models.GSLBPool{
		ID:        len(gslbPools) + 1,
		PolicyID:  policyID,
		Name:      req.Name,
		Priority:  req.Priority,
		Weight:    req.Weight,
		Enabled:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if pool.Priority == 0 {
		pool.Priority = 1
	}
	if pool.Weight == 0 {
		pool.Weight = 100
	}

	gslbPools = append(gslbPools, pool)
	c.JSON(http.StatusCreated, pool)
}

func GetGSLBPool(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pool ID"})
		return
	}

	for _, pool := range gslbPools {
		if pool.ID == id {
			c.JSON(http.StatusOK, pool)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Pool not found"})
}

func UpdateGSLBPool(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pool ID"})
		return
	}

	var req models.CreatePoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, pool := range gslbPools {
		if pool.ID == id {
			gslbPools[i].Name = req.Name
			gslbPools[i].Priority = req.Priority
			gslbPools[i].Weight = req.Weight
			gslbPools[i].UpdatedAt = time.Now()
			c.JSON(http.StatusOK, gslbPools[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Pool not found"})
}

func DeleteGSLBPool(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pool ID"})
		return
	}

	for i, pool := range gslbPools {
		if pool.ID == id {
			gslbPools = append(gslbPools[:i], gslbPools[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Pool deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Pool not found"})
}

func GetGSLBServers(c *gin.Context) {
	poolID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pool ID"})
		return
	}

	var servers []models.GSLBServer
	for _, server := range gslbServers {
		if server.PoolID == poolID {
			servers = append(servers, server)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"servers": servers,
		"total":   len(servers),
	})
}

func CreateGSLBServer(c *gin.Context) {
	poolID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pool ID"})
		return
	}

	var req models.CreateServerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	server := models.GSLBServer{
		ID:          len(gslbServers) + 1,
		PoolID:      poolID,
		Name:        req.Name,
		IPAddress:   req.IPAddress,
		Port:        req.Port,
		Weight:      req.Weight,
		Priority:    req.Priority,
		HealthCheck: req.HealthCheck,
		Status:      "unknown",
		Enabled:     true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if server.Port == 0 {
		server.Port = 80
	}
	if server.Weight == 0 {
		server.Weight = 100
	}
	if server.Priority == 0 {
		server.Priority = 1
	}
	if server.HealthCheck == "" {
		server.HealthCheck = "tcp"
	}

	gslbServers = append(gslbServers, server)
	c.JSON(http.StatusCreated, server)
}

func GetGSLBServer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid server ID"})
		return
	}

	for _, server := range gslbServers {
		if server.ID == id {
			c.JSON(http.StatusOK, server)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
}

func UpdateGSLBServer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid server ID"})
		return
	}

	var req models.CreateServerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, server := range gslbServers {
		if server.ID == id {
			gslbServers[i].Name = req.Name
			gslbServers[i].IPAddress = req.IPAddress
			gslbServers[i].Port = req.Port
			gslbServers[i].Weight = req.Weight
			gslbServers[i].Priority = req.Priority
			gslbServers[i].HealthCheck = req.HealthCheck
			gslbServers[i].UpdatedAt = time.Now()
			c.JSON(http.StatusOK, gslbServers[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
}

func DeleteGSLBServer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid server ID"})
		return
	}

	for i, server := range gslbServers {
		if server.ID == id {
			gslbServers = append(gslbServers[:i], gslbServers[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Server deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
}
