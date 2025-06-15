#!/bin/bash

set -e

echo "=== DNS/DHCP/GSLB Management System Deployment Script ==="

detect_os() {
    if [ -f /etc/os-release ]; then
        . /etc/os-release
        OS=$NAME
        VER=$VERSION_ID
    elif type lsb_release >/dev/null 2>&1; then
        OS=$(lsb_release -si)
        VER=$(lsb_release -sr)
    else
        echo "Cannot detect OS. Exiting."
        exit 1
    fi
    echo "Detected OS: $OS $VER"
}

install_packages() {
    case "$OS" in
        *"Ubuntu"*|*"Debian"*)
            echo "Installing packages with apt..."
            apt-get update
            apt-get install -y bind9 bind9utils bind9-doc isc-dhcp-server
            ;;
        *"CentOS"*|*"Red Hat"*|*"Rocky"*|*"AlmaLinux"*)
            echo "Installing packages with yum/dnf..."
            if command -v dnf &> /dev/null; then
                dnf install -y bind bind-utils dhcp-server
            else
                yum install -y bind bind-utils dhcp-server
            fi
            ;;
        *"SUSE"*|*"openSUSE"*)
            echo "Installing packages with zypper..."
            zypper install -y bind bind-utils dhcp-server
            ;;
        *)
            echo "Unsupported OS: $OS"
            exit 1
            ;;
    esac
}

setup_directories() {
    echo "Setting up directories..."
    mkdir -p /etc/bind/zones
    mkdir -p /var/lib/bind
    mkdir -p /var/log/bind
    mkdir -p /etc/dhcp
    mkdir -p /var/lib/dhcp
    
    if id "bind" &>/dev/null; then
        chown -R bind:bind /etc/bind /var/lib/bind /var/log/bind
    fi
    
    if id "dhcpd" &>/dev/null; then
        chown -R dhcpd:dhcpd /etc/dhcp /var/lib/dhcp
    elif id "_dhcp" &>/dev/null; then
        chown -R _dhcp:_dhcp /etc/dhcp /var/lib/dhcp
    fi
    
    chmod 755 /etc/bind /etc/bind/zones
    chmod 644 /etc/bind/named.conf 2>/dev/null || true
}

setup_services() {
    echo "Configuring systemd services..."
    systemctl enable bind9 2>/dev/null || systemctl enable named 2>/dev/null || true
    systemctl enable isc-dhcp-server 2>/dev/null || systemctl enable dhcpd 2>/dev/null || true
    
    echo "Services configured for automatic startup"
}

create_default_configs() {
    echo "Creating default configurations..."
    
    cat > /etc/bind/named.conf.local << 'EOF'
// Local zones configuration
// This file will be managed by the DNS/DHCP/GSLB Management System
EOF

    cat > /etc/dhcp/dhcpd.conf << 'EOF'
default-lease-time 600;
max-lease-time 7200;
authoritative;

option domain-name-servers 8.8.8.8, 8.8.4.4;
EOF

    touch /var/lib/dhcp/dhcpd.leases
    chmod 644 /var/lib/dhcp/dhcpd.leases
}

main() {
    if [ "$EUID" -ne 0 ]; then
        echo "Please run as root (use sudo)"
        exit 1
    fi
    
    detect_os
    install_packages
    setup_directories
    setup_services
    create_default_configs
    
    echo "=== Deployment completed successfully ==="
    echo "Next steps:"
    echo "1. Start the Go backend: cd backend && go run cmd/server/main.go"
    echo "2. Start the frontend: cd frontend && npm run dev"
    echo "3. Access the web interface at http://localhost:5173"
    echo "4. Use the web interface to configure DNS zones and DHCP scopes"
    echo "5. Deploy configurations using the 'Deploy Configuration' button"
}

main "$@"
