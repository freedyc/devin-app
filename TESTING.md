# Testing Guide for DNS/DHCP/GSLB Management System

## Overview

This guide provides comprehensive testing procedures for the DNS/DHCP/GSLB Management System, covering deployment, functionality, and integration testing.

## Prerequisites

- Linux system (Ubuntu/Debian, CentOS/RHEL, or SUSE)
- Root access for deployment testing
- Network connectivity for service testing
- Go 1.19+ and Node.js 16+ for development testing

## Testing Procedures

### 1. Deployment Script Testing

#### Test the deployment script functionality:

```bash
# Make the test script executable
chmod +x test_deployment.sh

# Run the comprehensive test suite
./test_deployment.sh
```

#### Manual deployment testing:

```bash
# Test the deployment script
sudo ./scripts/deploy.sh

# Verify services are installed
systemctl status bind9
systemctl status isc-dhcp-server

# Check configuration files
ls -la /etc/bind/
ls -la /etc/dhcp/
```

### 2. Backend API Testing

#### Start the backend server:

```bash
cd backend
go run cmd/server/main.go
```

#### Test basic API endpoints:

```bash
# Health check
curl http://localhost:8080/api/health

# DNS zones
curl http://localhost:8080/api/dns/zones

# Enhanced DNS features
curl http://localhost:8080/api/dns/enhanced/views
curl http://localhost:8080/api/dns/enhanced/forwarding-zones
curl http://localhost:8080/api/dns/enhanced/stub-zones
curl http://localhost:8080/api/dns/enhanced/recursion

# DHCP scopes
curl http://localhost:8080/api/dhcp/scopes

# Cluster nodes
curl http://localhost:8080/api/cluster/nodes
```

### 3. Frontend Testing

#### Start the frontend development server:

```bash
cd frontend
npm install
npm run dev
```

#### Access the web interface:

- Open http://localhost:5173 in your browser
- Test navigation between different sections
- Verify all pages load correctly
- Test responsive design on different screen sizes

### 4. Enhanced DNS Features Testing

#### Test DNS Views:

```bash
# Create a DNS view via API
curl -X POST http://localhost:8080/api/dns/enhanced/views \
  -H "Content-Type: application/json" \
  -d '{
    "name": "internal",
    "match_clients": "192.168.1.0/24",
    "recursion": true
  }'

# Verify view creation
curl http://localhost:8080/api/dns/enhanced/views
```

#### Test Forwarding Zones:

```bash
# Create a forwarding zone
curl -X POST http://localhost:8080/api/dns/enhanced/forwarding-zones \
  -H "Content-Type: application/json" \
  -d '{
    "name": "example.com",
    "forwarders": ["8.8.8.8", "8.8.4.4"],
    "forward": "first"
  }'

# Verify forwarding zone
curl http://localhost:8080/api/dns/enhanced/forwarding-zones
```

#### Test Stub Zones:

```bash
# Create a stub zone
curl -X POST http://localhost:8080/api/dns/enhanced/stub-zones \
  -H "Content-Type: application/json" \
  -d '{
    "name": "test.local",
    "masters": ["192.168.1.10", "192.168.1.11"]
  }'

# Verify stub zone
curl http://localhost:8080/api/dns/enhanced/stub-zones
```

### 5. Configuration Deployment Testing

#### Test DNS configuration deployment:

```bash
# Deploy DNS configuration
curl -X POST http://localhost:8080/api/dns/enhanced/deploy

# Verify BIND9 configuration
sudo named-checkconf /etc/bind/named.conf

# Check BIND9 service status
sudo systemctl status bind9
```

#### Test DHCP configuration deployment:

```bash
# Deploy DHCP configuration
curl -X POST http://localhost:8080/api/deploy/dhcp

# Verify DHCP configuration
sudo dhcpd -t -cf /etc/dhcp/dhcpd.conf

# Check DHCP service status
sudo systemctl status isc-dhcp-server
```

### 6. Service Integration Testing

#### Test DNS resolution:

```bash
# Test DNS queries (after deploying configuration)
nslookup example.com localhost
dig @localhost example.com

# Test recursive queries
dig @localhost google.com
```

#### Test DHCP functionality:

```bash
# Check DHCP lease file
sudo cat /var/lib/dhcp/dhcpd.leases

# Test DHCP discover (requires network setup)
sudo dhclient -v eth0
```

### 7. Cluster Management Testing

#### Test cluster node management:

```bash
# Add a cluster node
curl -X POST http://localhost:8080/api/cluster/nodes \
  -H "Content-Type: application/json" \
  -d '{
    "name": "secondary-node",
    "ip_address": "192.168.1.20",
    "role": "secondary"
  }'

# Sync cluster configuration
curl -X POST http://localhost:8080/api/cluster/sync

# Verify cluster status
curl http://localhost:8080/api/cluster/nodes
```

### 8. End-to-End Testing

#### Complete workflow test:

1. Deploy the system using the deployment script
2. Start backend and frontend services
3. Create DNS zones and records via web interface
4. Configure enhanced DNS features (views, forwarding zones)
5. Set up DHCP scopes and reservations
6. Deploy configurations to services
7. Test DNS resolution and DHCP functionality
8. Verify cluster synchronization

### 9. Performance Testing

#### Load testing:

```bash
# Test API performance
ab -n 1000 -c 10 http://localhost:8080/api/health

# Test DNS query performance
dig @localhost example.com +stats
```

### 10. Security Testing

#### Basic security checks:

```bash
# Check service permissions
ps aux | grep bind
ps aux | grep dhcp

# Verify configuration file permissions
ls -la /etc/bind/named.conf
ls -la /etc/dhcp/dhcpd.conf

# Test ACL functionality
dig @localhost example.com -b 192.168.1.100
```

## Expected Results

### Successful Deployment
- All packages installed correctly
- Services enabled and running
- Configuration files created with proper permissions
- Web interface accessible and functional

### API Functionality
- All endpoints respond correctly
- CRUD operations work for all resources
- Configuration deployment succeeds
- Error handling works properly

### Service Integration
- DNS queries resolve correctly
- DHCP leases are assigned properly
- Configuration changes take effect immediately
- Services restart/reload without errors

### Enhanced Features
- DNS views work with ACL matching
- Forwarding zones redirect queries correctly
- Stub zones delegate properly
- Recursion settings are applied

## Troubleshooting

### Common Issues

1. **Permission Errors**
   - Ensure proper file ownership and permissions
   - Run deployment script with sudo

2. **Service Start Failures**
   - Check configuration syntax
   - Review service logs: `journalctl -u bind9` or `journalctl -u isc-dhcp-server`

3. **API Connection Issues**
   - Verify backend is running on correct port
   - Check firewall settings

4. **Frontend Loading Issues**
   - Ensure all dependencies are installed
   - Check browser console for errors

### Log Locations

- Backend logs: Console output or configured log file
- BIND9 logs: `/var/log/bind/` or `journalctl -u bind9`
- DHCP logs: `/var/log/dhcp.log` or `journalctl -u isc-dhcp-server`
- System logs: `/var/log/syslog` or `journalctl`

This comprehensive testing guide ensures all components of the DNS/DHCP/GSLB Management System work correctly in both development and production environments.
