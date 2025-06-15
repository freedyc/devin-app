# DNS/DHCP/GSLB Management System

A comprehensive network management system similar to Infoblox, providing centralized management of DNS zones, DHCP scopes, and Global Server Load Balancing (GSLB) policies.

## Features

### DNS Management
- DNS zone management (create, edit, delete zones)
- DNS record management (A, AAAA, CNAME, MX, TXT, PTR records)
- BIND9 integration with configuration generation
- Real-time configuration validation

### DHCP Management
- DHCP scope management with IP range configuration
- Static IP reservations by MAC address
- Active lease monitoring
- ISC DHCP server integration

### GSLB Management
- Load balancing policy configuration
- Server pool management with health checks
- Weighted and priority-based load balancing
- Geographic load balancing support

### Configuration Management
- Service configuration validation
- Safe configuration reload/restart
- Backup and restore functionality
- Audit logging

## Architecture

```
├── frontend/          # Vite + React + TypeScript web interface
├── backend/           # Go API server with Gin framework
├── configs/           # Sample BIND9 and DHCP configurations
├── scripts/           # Setup and utility scripts
└── docker-compose.yml # Development environment
```

## Technology Stack

- **Frontend**: Vite, React, TypeScript, Tailwind CSS
- **Backend**: Go, Gin web framework, PostgreSQL
- **Services**: BIND9 DNS server, ISC DHCP server
- **Development**: Docker Compose for local services

## Quick Start

### Prerequisites
- Docker and Docker Compose
- Go 1.21+
- Node.js 18+
- Git

### Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd devin-app
```

2. Run the setup script:
```bash
chmod +x scripts/setup.sh
./scripts/setup.sh
```

3. Start the development environment:
```bash
# Start supporting services
docker-compose up -d

# Start the Go backend (in a new terminal)
cd backend
go run cmd/server/main.go

# Start the frontend (in another terminal)
cd frontend
npm run dev
```

4. Access the application:
- Web Interface: http://localhost:5173
- API Server: http://localhost:8080
- API Health Check: http://localhost:8080/api/health

## Development

### Backend Development

The Go backend provides REST APIs for managing DNS, DHCP, and GSLB configurations:

```bash
cd backend
go run cmd/server/main.go
```

Key API endpoints:
- `/api/dns/zones` - DNS zone management
- `/api/dhcp/scopes` - DHCP scope management
- `/api/gslb/policies` - GSLB policy management
- `/api/config/*` - Configuration management

### Frontend Development

The React frontend provides a modern web interface:

```bash
cd frontend
npm run dev
```

Features:
- Responsive design with Tailwind CSS
- Real-time data updates
- Form validation and error handling
- Modern React patterns with hooks

### Configuration Integration

The system integrates with BIND9 and ISC DHCP server through:
- Configuration file generation from API data
- Syntax validation before applying changes
- Safe service reload/restart mechanisms
- Configuration backup and rollback

## API Documentation

### DNS Management
- `GET /api/dns/zones` - List all DNS zones
- `POST /api/dns/zones` - Create a new DNS zone
- `GET /api/dns/zones/{id}/records` - Get records for a zone
- `POST /api/dns/zones/{id}/records` - Create a new DNS record

### DHCP Management
- `GET /api/dhcp/scopes` - List all DHCP scopes
- `POST /api/dhcp/scopes` - Create a new DHCP scope
- `GET /api/dhcp/leases` - List active DHCP leases
- `POST /api/dhcp/scopes/{id}/reservations` - Create IP reservation

### GSLB Management
- `GET /api/gslb/policies` - List all GSLB policies
- `POST /api/gslb/policies` - Create a new GSLB policy
- `GET /api/gslb/policies/{id}/pools` - Get pools for a policy
- `POST /api/gslb/pools/{id}/servers` - Add server to pool

## Configuration

### Environment Variables

Backend configuration:
- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_USER` - Database user (default: admin)
- `DB_PASSWORD` - Database password (default: password)
- `BIND9_CONFIG_PATH` - BIND9 config file path
- `DHCP_CONFIG_PATH` - DHCP config file path

Frontend configuration:
- `VITE_API_URL` - Backend API URL (default: http://localhost:8080)

### Service Integration

The system manages BIND9 and DHCP configurations through:
- Direct configuration file manipulation
- Template-based configuration generation
- Service validation and reload commands
- Error handling and rollback mechanisms

## Security

- Input validation and sanitization
- Configuration file permission management
- Audit logging for all administrative actions
- Safe configuration update procedures

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## License

This project is open source and available under the MIT License.
