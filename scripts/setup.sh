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
