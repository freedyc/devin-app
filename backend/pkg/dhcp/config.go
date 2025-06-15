package dhcp

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/freedyc/devin-app/backend/internal/models"
)

const dhcpdConfTemplate = `# DHCP Server Configuration
# Generated automatically - do not edit manually

default-lease-time 600;
max-lease-time 7200;
authoritative;

option domain-name-servers 8.8.8.8, 8.8.4.4;

{{range .Scopes}}{{if .Enabled}}
# Scope: {{.Name}}
subnet {{.Network}} netmask {{.Netmask}} {
	range {{.RangeStart}} {{.RangeEnd}};
	option routers {{.Gateway}};
	{{if .DNSServers}}option domain-name-servers {{.DNSServers}};{{end}}
	default-lease-time {{.LeaseTime}};
	max-lease-time {{.LeaseTime}};
}

{{range $.Reservations}}{{if and .Enabled (eq .ScopeID $.ID)}}
# Reservation: {{.Hostname}}
host {{.Hostname}} {
	hardware ethernet {{.MACAddress}};
	fixed-address {{.IPAddress}};
}
{{end}}{{end}}
{{end}}{{end}}`

type DHCPConfig struct {
	Scopes       []models.DHCPScope
	Reservations []models.DHCPReservation
}

func GenerateDHCPConf(scopes []models.DHCPScope, reservations []models.DHCPReservation) (string, error) {
	tmpl, err := template.New("dhcpd").Parse(dhcpdConfTemplate)
	if err != nil {
		return "", err
	}

	config := DHCPConfig{
		Scopes:       scopes,
		Reservations: reservations,
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, config)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func WriteDHCPConf(configPath string, scopes []models.DHCPScope, reservations []models.DHCPReservation) error {
	content, err := GenerateDHCPConf(scopes, reservations)
	if err != nil {
		return err
	}

	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(configPath, []byte(content), 0644)
}

func ParseLeaseFile(leasesPath string) ([]models.DHCPLease, error) {
	content, err := os.ReadFile(leasesPath)
	if err != nil {
		return nil, err
	}

	var leases []models.DHCPLease
	lines := strings.Split(string(content), "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "lease ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				ip := parts[1]
				lease := models.DHCPLease{
					IPAddress: ip,
					State:     "active",
				}
				leases = append(leases, lease)
			}
		}
	}

	return leases, nil
}
