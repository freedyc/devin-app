import { useEffect, useState } from 'react'
import { Plus, Trash2, Globe } from 'lucide-react'
import { dnsApi, type DNSZone, type DNSRecord } from '@/services/api'

export default function DNSManagement() {
  const [zones, setZones] = useState<DNSZone[]>([])
  const [selectedZone, setSelectedZone] = useState<DNSZone | null>(null)
  const [records, setRecords] = useState<DNSRecord[]>([])
  const [loading, setLoading] = useState(true)


  useEffect(() => {
    fetchZones()
  }, [])

  const fetchZones = async () => {
    try {
      const response = await dnsApi.getZones()
      setZones(response.data.zones)
    } catch (error) {
      console.error('Failed to fetch zones:', error)
    } finally {
      setLoading(false)
    }
  }

  const fetchRecords = async (zoneId: number) => {
    try {
      const response = await dnsApi.getRecords(zoneId)
      setRecords(response.data.records)
    } catch (error) {
      console.error('Failed to fetch records:', error)
    }
  }

  const handleZoneSelect = (zone: DNSZone) => {
    setSelectedZone(zone)
    fetchRecords(zone.id)
  }

  const handleDeleteZone = async (id: number) => {
    if (confirm('Are you sure you want to delete this zone?')) {
      try {
        await dnsApi.deleteZone(id)
        fetchZones()
        if (selectedZone?.id === id) {
          setSelectedZone(null)
          setRecords([])
        }
      } catch (error) {
        console.error('Failed to delete zone:', error)
      }
    }
  }

  const handleDeleteRecord = async (id: number) => {
    if (confirm('Are you sure you want to delete this record?')) {
      try {
        await dnsApi.deleteRecord(id)
        if (selectedZone) {
          fetchRecords(selectedZone.id)
        }
      } catch (error) {
        console.error('Failed to delete record:', error)
      }
    }
  }

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      </div>
    )
  }

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-gray-900">DNS Management</h1>
        <p className="mt-1 text-sm text-gray-600">
          Manage DNS zones and records
        </p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                DNS Zones
              </h3>
              <button
                className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
              >
                <Plus className="h-4 w-4 mr-1" />
                Add Zone
              </button>
            </div>

            <div className="space-y-3">
              {zones.map((zone) => (
                <div
                  key={zone.id}
                  className={`p-3 border rounded-lg cursor-pointer transition-colors ${
                    selectedZone?.id === zone.id
                      ? 'border-blue-500 bg-blue-50'
                      : 'border-gray-200 hover:border-gray-300'
                  }`}
                  onClick={() => handleZoneSelect(zone)}
                >
                  <div className="flex items-center justify-between">
                    <div className="flex items-center">
                      <Globe className="h-5 w-5 text-gray-400 mr-2" />
                      <div>
                        <p className="text-sm font-medium text-gray-900">
                          {zone.name}
                        </p>
                        <p className="text-xs text-gray-500">
                          {zone.type} zone
                        </p>
                      </div>
                    </div>
                    <div className="flex items-center space-x-2">
                      <span
                        className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                          zone.enabled
                            ? 'bg-green-100 text-green-800'
                            : 'bg-red-100 text-red-800'
                        }`}
                      >
                        {zone.enabled ? 'Enabled' : 'Disabled'}
                      </span>
                      <button
                        onClick={(e) => {
                          e.stopPropagation()
                          handleDeleteZone(zone.id)
                        }}
                        className="text-red-600 hover:text-red-900"
                      >
                        <Trash2 className="h-4 w-4" />
                      </button>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>

        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                DNS Records
                {selectedZone && (
                  <span className="text-sm font-normal text-gray-500 ml-2">
                    for {selectedZone.name}
                  </span>
                )}
              </h3>
              {selectedZone && (
                <button
                  className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-green-600 hover:bg-green-700"
                >
                  <Plus className="h-4 w-4 mr-1" />
                  Add Record
                </button>
              )}
            </div>

            {selectedZone ? (
              <div className="space-y-3">
                {records.map((record) => (
                  <div
                    key={record.id}
                    className="p-3 border border-gray-200 rounded-lg"
                  >
                    <div className="flex items-center justify-between">
                      <div>
                        <div className="flex items-center space-x-2">
                          <span className="text-sm font-medium text-gray-900">
                            {record.name}
                          </span>
                          <span className="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-gray-100 text-gray-800">
                            {record.type}
                          </span>
                        </div>
                        <p className="text-sm text-gray-600 mt-1">
                          {record.value}
                        </p>
                        <p className="text-xs text-gray-500">
                          TTL: {record.ttl}s
                        </p>
                      </div>
                      <div className="flex items-center space-x-2">
                        <span
                          className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                            record.enabled
                              ? 'bg-green-100 text-green-800'
                              : 'bg-red-100 text-red-800'
                          }`}
                        >
                          {record.enabled ? 'Enabled' : 'Disabled'}
                        </span>
                        <button
                          onClick={() => handleDeleteRecord(record.id)}
                          className="text-red-600 hover:text-red-900"
                        >
                          <Trash2 className="h-4 w-4" />
                        </button>
                      </div>
                    </div>
                  </div>
                ))}
                {records.length === 0 && (
                  <p className="text-gray-500 text-center py-4">
                    No records found for this zone
                  </p>
                )}
              </div>
            ) : (
              <p className="text-gray-500 text-center py-8">
                Select a zone to view its records
              </p>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
