#!/bin/bash

echo "Setting up DNS/DHCP/GSLB Management System..."

mkdir -p data/bind9
mkdir -p data/dhcp

chmod 755 data/bind9
chmod 755 data/dhcp

echo "Installing Go dependencies..."
cd backend
go mod tidy
cd ..

echo "Installing Node.js dependencies..."
cd frontend
npm install
cd ..

echo "Setup complete!"
echo ""
echo "To start the development environment:"
echo "1. Start Docker services: docker-compose up -d"
echo "2. Start Go backend: cd backend && go run cmd/server/main.go"
echo "3. Start frontend: cd frontend && npm run dev"
echo ""
echo "For production deployment:"
echo "1. Run deployment script: sudo ./scripts/deploy.sh"
echo "2. Configure firewall rules for DNS (port 53) and DHCP (ports 67/68)"
echo "3. Update /etc/bind/named.conf and /etc/dhcp/dhcpd.conf as needed"
echo "4. Start services: sudo systemctl start bind9 isc-dhcp-server"
echo ""
echo "For cluster deployment:"
echo "1. Run deployment script on all nodes"
echo "2. Configure primary node first, then secondary nodes"
echo "3. Use the web interface to manage cluster synchronization"
