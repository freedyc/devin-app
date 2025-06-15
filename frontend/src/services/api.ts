import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'

const api = axios.create({
  baseURL: `${API_BASE_URL}/api`,
  headers: {
    'Content-Type': 'application/json',
  },
})

export interface DNSZone {
  id: number
  name: string
  type: string
  serial: number
  refresh: number
  retry: number
  expire: number
  min_ttl: number
  primary_ns: string
  admin_email: string
  enabled: boolean
  created_at: string
  updated_at: string
}

export interface DNSRecord {
  id: number
  zone_id: number
  name: string
  type: string
  value: string
  ttl: number
  priority?: number
  weight?: number
  port?: number
  enabled: boolean
  created_at: string
  updated_at: string
}

export interface DHCPScope {
  id: number
  name: string
  network: string
  netmask: string
  range_start: string
  range_end: string
  gateway: string
  dns_servers: string
  lease_time: number
  enabled: boolean
  created_at: string
  updated_at: string
}

export interface DHCPReservation {
  id: number
  scope_id: number
  mac_address: string
  ip_address: string
  hostname: string
  description: string
  enabled: boolean
  created_at: string
  updated_at: string
}

export interface DHCPLease {
  id: number
  ip_address: string
  mac_address: string
  hostname: string
  start_time: string
  end_time: string
  state: string
  scope_id: number
}

export interface GSLBPolicy {
  id: number
  name: string
  fqdn: string
  method: string
  ttl: number
  enabled: boolean
  description: string
  created_at: string
  updated_at: string
}

export interface GSLBPool {
  id: number
  policy_id: number
  name: string
  priority: number
  weight: number
  enabled: boolean
  created_at: string
  updated_at: string
}

export interface GSLBServer {
  id: number
  pool_id: number
  name: string
  ip_address: string
  port: number
  weight: number
  priority: number
  health_check: string
  status: string
  enabled: boolean
  created_at: string
  updated_at: string
}

export const dnsApi = {
  getZones: () => api.get<{ zones: DNSZone[]; total: number }>('/dns/zones'),
  createZone: (data: Partial<DNSZone>) => api.post<DNSZone>('/dns/zones', data),
  getZone: (id: number) => api.get<DNSZone>(`/dns/zones/${id}`),
  updateZone: (id: number, data: Partial<DNSZone>) => api.put<DNSZone>(`/dns/zones/${id}`, data),
  deleteZone: (id: number) => api.delete(`/dns/zones/${id}`),
  getRecords: (zoneId: number) => api.get<{ records: DNSRecord[]; total: number }>(`/dns/zones/${zoneId}/records`),
  createRecord: (zoneId: number, data: Partial<DNSRecord>) => api.post<DNSRecord>(`/dns/zones/${zoneId}/records`, data),
  getRecord: (id: number) => api.get<DNSRecord>(`/dns/records/${id}`),
  updateRecord: (id: number, data: Partial<DNSRecord>) => api.put<DNSRecord>(`/dns/records/${id}`, data),
  deleteRecord: (id: number) => api.delete(`/dns/records/${id}`),
}

export const dhcpApi = {
  getScopes: () => api.get<{ scopes: DHCPScope[]; total: number }>('/dhcp/scopes'),
  createScope: (data: Partial<DHCPScope>) => api.post<DHCPScope>('/dhcp/scopes', data),
  getScope: (id: number) => api.get<DHCPScope>(`/dhcp/scopes/${id}`),
  updateScope: (id: number, data: Partial<DHCPScope>) => api.put<DHCPScope>(`/dhcp/scopes/${id}`, data),
  deleteScope: (id: number) => api.delete(`/dhcp/scopes/${id}`),
  getReservations: (scopeId: number) => api.get<{ reservations: DHCPReservation[]; total: number }>(`/dhcp/scopes/${scopeId}/reservations`),
  createReservation: (scopeId: number, data: Partial<DHCPReservation>) => api.post<DHCPReservation>(`/dhcp/scopes/${scopeId}/reservations`, data),
  getReservation: (id: number) => api.get<DHCPReservation>(`/dhcp/reservations/${id}`),
  updateReservation: (id: number, data: Partial<DHCPReservation>) => api.put<DHCPReservation>(`/dhcp/reservations/${id}`, data),
  deleteReservation: (id: number) => api.delete(`/dhcp/reservations/${id}`),
  getLeases: () => api.get<{ leases: DHCPLease[]; total: number }>('/dhcp/leases'),
  getLeasesByScope: (scopeId: number) => api.get<{ leases: DHCPLease[]; total: number }>(`/dhcp/scopes/${scopeId}/leases`),
}

export const gslbApi = {
  getPolicies: () => api.get<{ policies: GSLBPolicy[]; total: number }>('/gslb/policies'),
  createPolicy: (data: Partial<GSLBPolicy>) => api.post<GSLBPolicy>('/gslb/policies', data),
  getPolicy: (id: number) => api.get<GSLBPolicy>(`/gslb/policies/${id}`),
  updatePolicy: (id: number, data: Partial<GSLBPolicy>) => api.put<GSLBPolicy>(`/gslb/policies/${id}`, data),
  deletePolicy: (id: number) => api.delete(`/gslb/policies/${id}`),
  getPools: (policyId: number) => api.get<{ pools: GSLBPool[]; total: number }>(`/gslb/policies/${policyId}/pools`),
  createPool: (policyId: number, data: Partial<GSLBPool>) => api.post<GSLBPool>(`/gslb/policies/${policyId}/pools`, data),
  getPool: (id: number) => api.get<GSLBPool>(`/gslb/pools/${id}`),
  updatePool: (id: number, data: Partial<GSLBPool>) => api.put<GSLBPool>(`/gslb/pools/${id}`, data),
  deletePool: (id: number) => api.delete(`/gslb/pools/${id}`),
  getServers: (poolId: number) => api.get<{ servers: GSLBServer[]; total: number }>(`/gslb/pools/${poolId}/servers`),
  createServer: (poolId: number, data: Partial<GSLBServer>) => api.post<GSLBServer>(`/gslb/pools/${poolId}/servers`, data),
  getServer: (id: number) => api.get<GSLBServer>(`/gslb/servers/${id}`),
  updateServer: (id: number, data: Partial<GSLBServer>) => api.put<GSLBServer>(`/gslb/servers/${id}`, data),
  deleteServer: (id: number) => api.delete(`/gslb/servers/${id}`),
}

export const configApi = {
  reloadBind9: () => api.post('/config/bind9/reload'),
  restartDHCP: () => api.post('/config/dhcp/restart'),
  validateBind9: () => api.get('/config/bind9/validate'),
  validateDHCP: () => api.get('/config/dhcp/validate'),
}
