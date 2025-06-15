#!/bin/bash

set -e

echo "=== Testing Enhanced DNS Functionality ==="

echo "Starting backend server..."
cd backend
./dns-dhcp-server &
BACKEND_PID=$!
cd ..

sleep 3

echo "Testing enhanced DNS API endpoints..."

echo "Creating DNS view..."
curl -X POST http://localhost:8080/api/dns/enhanced/views \
  -H "Content-Type: application/json" \
  -d '{
    "name": "internal-view",
    "match_clients": "192.168.0.0/16",
    "recursion": true
  }'

echo -e "\n\nListing DNS views..."
curl http://localhost:8080/api/dns/enhanced/views

echo -e "\n\nCreating forwarding zone..."
curl -X POST http://localhost:8080/api/dns/enhanced/forwarding-zones \
  -H "Content-Type: application/json" \
  -d '{
    "name": "external.com",
    "forwarders": ["1.1.1.1", "1.0.0.1"],
    "forward": "only"
  }'

echo -e "\n\nListing forwarding zones..."
curl http://localhost:8080/api/dns/enhanced/forwarding-zones

echo -e "\n\nCreating stub zone..."
curl -X POST http://localhost:8080/api/dns/enhanced/stub-zones \
  -H "Content-Type: application/json" \
  -d '{
    "name": "test.local",
    "masters": ["192.168.1.10", "192.168.1.11"]
  }'

echo -e "\n\nListing stub zones..."
curl http://localhost:8080/api/dns/enhanced/stub-zones

echo -e "\n\nGetting recursion config..."
curl http://localhost:8080/api/dns/enhanced/recursion

echo -e "\n\nListing cluster nodes..."
curl http://localhost:8080/api/cluster/nodes

echo -e "\n\nCreating cluster node..."
curl -X POST http://localhost:8080/api/cluster/nodes \
  -H "Content-Type: application/json" \
  -d '{
    "name": "secondary-node",
    "ip_address": "192.168.1.20",
    "role": "secondary"
  }'

echo -e "\n\nTesting cluster sync..."
curl -X POST http://localhost:8080/api/cluster/sync

echo -e "\n\nTesting DNS configuration deployment..."
curl -X POST http://localhost:8080/api/dns/enhanced/deploy

echo -e "\n\nCleaning up..."
kill $BACKEND_PID 2>/dev/null || true

echo -e "\n=== Enhanced DNS functionality test completed ==="
