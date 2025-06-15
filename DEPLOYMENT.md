# Deployment Guide

This guide covers deploying the DNS/DHCP/GSLB Management System in production environments.

## Prerequisites

- Linux server (Ubuntu/Debian, CentOS/RHEL, or SUSE)
- Root access for service installation
- Go 1.19+ for backend compilation
- Node.js 16+ for frontend building

## One-Click Deployment

The system includes a comprehensive deployment script that automatically installs and configures bind9 and dhcpd services:

```bash
sudo ./scripts/deploy.sh
```

This script will:
- Detect your operating system
- Install bind9 and dhcpd packages using the appropriate package manager
- Create necessary directories with proper permissions
- Set up default configuration files
- Configure systemd services for automatic startup

## Manual Deployment Steps

If you prefer manual deployment or need to customize the process:

### 1. Install System Dependencies

**Ubuntu/Debian:**
```bash
sudo apt-get update
sudo apt-get install -y bind9 bind9utils bind9-doc isc-dhcp-server
```

**CentOS/RHEL:**
```bash
sudo dnf install -y bind bind-utils dhcp-server
# or for older versions:
sudo yum install -y bind bind-utils dhcp-server
```

### 2. Set Up Directory Structure

```bash
sudo mkdir -p /etc/bind/zones
sudo mkdir -p /var/lib/bind
sudo mkdir -p /var/log/bind
sudo mkdir -p /etc/dhcp
sudo mkdir -p /var/lib/dhcp

sudo chown -R bind:bind /etc/bind /var/lib/bind /var/log/bind
sudo chown -R dhcpd:dhcpd /etc/dhcp /var/lib/dhcp
```

### 3. Build and Deploy Application

```bash
# Build backend
cd backend
go build -o dns-dhcp-server cmd/server/main.go

# Build frontend
cd ../frontend
npm install
npm run build

# Deploy frontend files to web server
sudo cp -r dist/* /var/www/html/
```

### 4. Configure Services

```bash
# Enable services
sudo systemctl enable bind9
sudo systemctl enable isc-dhcp-server

# Start services
sudo systemctl start bind9
sudo systemctl start isc-dhcp-server
```

## Cluster Deployment

For high availability deployments:

### Primary Node Setup

1. Run the deployment script on the primary node
2. Configure it as the primary in the web interface
3. Set up DNS zone transfers and DHCP failover

### Secondary Node Setup

1. Run the deployment script on secondary nodes
2. Add them to the cluster via the web interface
3. Use the "Sync Configuration" feature to distribute settings

## Configuration Management

The system provides real-time configuration deployment:

1. **DNS Configuration**: Changes made in the web interface can be deployed to bind9 using the "Deploy DNS Configuration" button
2. **DHCP Configuration**: DHCP scope changes are applied using the "Deploy DHCP Configuration" button
3. **System-wide Deployment**: Use "Deploy System Configuration" to apply all changes at once

## Security Considerations

- Run services with minimal privileges
- Configure firewall rules for DNS (port 53) and DHCP (ports 67/68)
- Use ACLs in DNS views to restrict access
- Regularly update system packages and dependencies

## Monitoring and Maintenance

- Monitor service logs: `journalctl -u bind9` and `journalctl -u isc-dhcp-server`
- Check configuration syntax before deployment
- Use the web interface dashboard for real-time status monitoring
- Set up automated backups of configuration files

## Troubleshooting

### Common Issues

1. **Permission Errors**: Ensure proper ownership of configuration directories
2. **Service Start Failures**: Check configuration syntax and log files
3. **Network Connectivity**: Verify firewall rules and network interfaces
4. **Cluster Sync Issues**: Check network connectivity between nodes

### Log Locations

- BIND9 logs: `/var/log/bind/` or via `journalctl -u bind9`
- DHCP logs: `/var/log/dhcp.log` or via `journalctl -u isc-dhcp-server`
- Application logs: Check backend console output

## Advanced Features

### DNS Views

Configure different DNS responses for different client networks:

1. Create views in the Enhanced DNS section
2. Define ACLs for client matching
3. Assign zones to specific views
4. Deploy configuration to activate

### Forwarding Zones

Set up conditional forwarding for specific domains:

1. Add forwarding zones in Enhanced DNS
2. Specify forwarder IP addresses
3. Choose forwarding mode (first or only)
4. Deploy to activate forwarding

### Stub Zones

Configure stub zones for delegation:

1. Create stub zones in Enhanced DNS
2. Specify master server IP addresses
3. Deploy configuration to activate delegation

This deployment guide ensures a robust, scalable DNS/DHCP management solution suitable for enterprise environments.
