package bind9

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/freedyc/devin-app/backend/internal/models"
)

type ZoneConfig struct {
	Zone    models.DNSZone
	Records []models.DNSRecord
}

const zoneTemplate = `$TTL {{.Zone.MinTTL}}
@	IN	SOA	{{.Zone.PrimaryNS}}. {{.Zone.AdminEmail}}. (
		{{.Zone.Serial}}	; Serial
		{{.Zone.Refresh}}	; Refresh
		{{.Zone.Retry}}		; Retry
		{{.Zone.Expire}}	; Expire
		{{.Zone.MinTTL}} )	; Negative Cache TTL

; Name servers
@	IN	NS	{{.Zone.PrimaryNS}}.

; Records
{{range .Records}}{{if .Enabled}}{{.Name}}	{{.TTL}}	IN	{{.Type}}	{{if eq .Type "MX"}}{{.Priority}} {{end}}{{.Value}}{{if and (eq .Type "SRV") .Weight .Port}}	{{.Weight}} {{.Port}}{{end}}
{{end}}{{end}}`

const namedConfTemplate = `options {
	directory "/var/lib/bind";
	recursion yes;
	allow-recursion { any; };
	listen-on { any; };
	listen-on-v6 { any; };
	forwarders {
		8.8.8.8;
		8.8.4.4;
	};
};

logging {
	channel default_debug {
		file "data/named.run";
		severity dynamic;
	};
};

{{range .}}{{if .Enabled}}zone "{{.Name}}" {
	type master;
	file "/etc/bind/zones/{{.Name}}.zone";
};

{{end}}{{end}}`

func GenerateZoneFile(zone models.DNSZone, records []models.DNSRecord) (string, error) {
	tmpl, err := template.New("zone").Parse(zoneTemplate)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, ZoneConfig{Zone: zone, Records: records})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func GenerateNamedConf(zones []models.DNSZone) (string, error) {
	tmpl, err := template.New("named").Parse(namedConfTemplate)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, zones)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func WriteZoneFile(zonesPath string, zone models.DNSZone, records []models.DNSRecord) error {
	content, err := GenerateZoneFile(zone, records)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(zonesPath, 0755); err != nil {
		return err
	}

	filename := filepath.Join(zonesPath, fmt.Sprintf("%s.zone", zone.Name))
	return os.WriteFile(filename, []byte(content), 0644)
}

func WriteNamedConf(configPath string, zones []models.DNSZone) error {
	content, err := GenerateNamedConf(zones)
	if err != nil {
		return err
	}

	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(configPath, []byte(content), 0644)
}

const enhancedNamedConfTemplate = `options {
	directory "/var/lib/bind";
	recursion {{if .Recursion.Enabled}}yes{{else}}no{{end}};
	{{if .Recursion.AllowRecursion}}allow-recursion { {{range .Recursion.AllowRecursion}}{{.}}; {{end}}};{{end}}
	listen-on { any; };
	listen-on-v6 { any; };
	{{if .Recursion.Forwarders}}forwarders {
		{{range .Recursion.Forwarders}}{{.}};
		{{end}}
	};{{end}}
	{{if .Recursion.ForwardFirst}}forward first;{{else}}forward only;{{end}}
	dnssec-validation {{.Recursion.DNSSECValidation}};
};

logging {
	channel default_debug {
		file "data/named.run";
		severity dynamic;
	};
};

{{range .Views}}
view "{{.Name}}" {
	match-clients { {{.MatchClients}}; };
	recursion {{if .Recursion}}yes{{else}}no{{end}};
	
	{{range $.Zones}}{{if and .Enabled (eq .ViewID $.ID)}}
	zone "{{.Name}}" {
		type {{.ZoneType}};
		{{if eq .ZoneType "master"}}file "/etc/bind/zones/{{.Name}}.zone";{{end}}
		{{if eq .ZoneType "slave"}}masters { {{.Masters}}; };{{end}}
		{{if eq .ZoneType "forward"}}forwarders { {{.Forwarders}}; };{{end}}
		{{if eq .ZoneType "stub"}}masters { {{.Masters}}; };{{end}}
	};
	{{end}}{{end}}
	
	{{range $.ForwardingZones}}{{if .Enabled}}
	zone "{{.Name}}" {
		type forward;
		forward {{.Forward}};
		forwarders { {{range .Forwarders}}{{.}}; {{end}}};
	};
	{{end}}{{end}}
	
	{{range $.StubZones}}{{if .Enabled}}
	zone "{{.Name}}" {
		type stub;
		masters { {{range .Masters}}{{.}}; {{end}}};
	};
	{{end}}{{end}}
};
{{end}}

{{if .DefaultZones}}
{{range .DefaultZones}}{{if .Enabled}}
zone "{{.Name}}" {
	type {{.ZoneType}};
	{{if eq .ZoneType "master"}}file "/etc/bind/zones/{{.Name}}.zone";{{end}}
	{{if eq .ZoneType "slave"}}masters { {{.Masters}}; };{{end}}
};
{{end}}{{end}}
{{end}}`

type EnhancedNamedConfig struct {
	Recursion       models.RecursionConfig
	Views           []models.DNSView
	Zones           []models.DNSZoneEnhanced
	ForwardingZones []models.ForwardingZone
	StubZones       []models.StubZone
	DefaultZones    []models.DNSZoneEnhanced
}

func GenerateEnhancedNamedConf(config EnhancedNamedConfig) (string, error) {
	tmpl, err := template.New("enhanced-named").Parse(enhancedNamedConfTemplate)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, config)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func WriteEnhancedNamedConf(configPath string, config EnhancedNamedConfig) error {
	content, err := GenerateEnhancedNamedConf(config)
	if err != nil {
		return err
	}

	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(configPath, []byte(content), 0644)
}
