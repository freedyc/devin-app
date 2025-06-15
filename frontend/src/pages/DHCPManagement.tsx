import { useEffect, useState } from 'react'
import { Plus, Edit, Trash2, Network, Monitor } from 'lucide-react'
import { dhcpApi, type DHCPScope, type DHCPReservation, type DHCPLease } from '@/services/api'

export default function DHCPManagement() {
  const [scopes, setScopes] = useState<DHCPScope[]>([])
  const [selectedScope, setSelectedScope] = useState<DHCPScope | null>(null)
  const [reservations, setReservations] = useState<DHCPReservation[]>([])
  const [leases, setLeases] = useState<DHCPLease[]>([])
  const [activeTab, setActiveTab] = useState<'reservations' | 'leases'>('reservations')
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchScopes()
  }, [])

  const fetchScopes = async () => {
    try {
      const response = await dhcpApi.getScopes()
      setScopes(response.data.scopes)
    } catch (error) {
      console.error('Failed to fetch scopes:', error)
    } finally {
      setLoading(false)
    }
  }

  const fetchReservations = async (scopeId: number) => {
    try {
      const response = await dhcpApi.getReservations(scopeId)
      setReservations(response.data.reservations)
    } catch (error) {
      console.error('Failed to fetch reservations:', error)
    }
  }

  const fetchLeases = async (scopeId: number) => {
    try {
      const response = await dhcpApi.getLeasesByScope(scopeId)
      setLeases(response.data.leases)
    } catch (error) {
      console.error('Failed to fetch leases:', error)
    }
  }

  const handleScopeSelect = (scope: DHCPScope) => {
    setSelectedScope(scope)
    fetchReservations(scope.id)
    fetchLeases(scope.id)
  }

  const handleDeleteScope = async (id: number) => {
    if (confirm('Are you sure you want to delete this scope?')) {
      try {
        await dhcpApi.deleteScope(id)
        fetchScopes()
        if (selectedScope?.id === id) {
          setSelectedScope(null)
          setReservations([])
          setLeases([])
        }
      } catch (error) {
        console.error('Failed to delete scope:', error)
      }
    }
  }

  const handleDeleteReservation = async (id: number) => {
    if (confirm('Are you sure you want to delete this reservation?')) {
      try {
        await dhcpApi.deleteReservation(id)
        if (selectedScope) {
          fetchReservations(selectedScope.id)
        }
      } catch (error) {
        console.error('Failed to delete reservation:', error)
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
        <h1 className="text-2xl font-bold text-gray-900">DHCP Management</h1>
        <p className="mt-1 text-sm text-gray-600">
          Manage DHCP scopes, reservations, and leases
        </p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                DHCP Scopes
              </h3>
              <button className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">
                <Plus className="h-4 w-4 mr-1" />
                Add Scope
              </button>
            </div>

            <div className="space-y-3">
              {scopes.map((scope) => (
                <div
                  key={scope.id}
                  className={`p-3 border rounded-lg cursor-pointer transition-colors ${
                    selectedScope?.id === scope.id
                      ? 'border-blue-500 bg-blue-50'
                      : 'border-gray-200 hover:border-gray-300'
                  }`}
                  onClick={() => handleScopeSelect(scope)}
                >
                  <div className="flex items-center justify-between">
                    <div className="flex items-center">
                      <Network className="h-5 w-5 text-gray-400 mr-2" />
                      <div>
                        <p className="text-sm font-medium text-gray-900">
                          {scope.name}
                        </p>
                        <p className="text-xs text-gray-500">
                          {scope.network}/{scope.netmask}
                        </p>
                        <p className="text-xs text-gray-500">
                          Range: {scope.range_start} - {scope.range_end}
                        </p>
                      </div>
                    </div>
                    <div className="flex items-center space-x-2">
                      <span
                        className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                          scope.enabled
                            ? 'bg-green-100 text-green-800'
                            : 'bg-red-100 text-red-800'
                        }`}
                      >
                        {scope.enabled ? 'Enabled' : 'Disabled'}
                      </span>
                      <button
                        onClick={(e) => {
                          e.stopPropagation()
                          handleDeleteScope(scope.id)
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
                {selectedScope ? selectedScope.name : 'Select a scope'}
              </h3>
              {selectedScope && (
                <button className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-green-600 hover:bg-green-700">
                  <Plus className="h-4 w-4 mr-1" />
                  Add {activeTab === 'reservations' ? 'Reservation' : 'Lease'}
                </button>
              )}
            </div>

            {selectedScope ? (
              <>
                <div className="border-b border-gray-200 mb-4">
                  <nav className="-mb-px flex space-x-8">
                    <button
                      onClick={() => setActiveTab('reservations')}
                      className={`py-2 px-1 border-b-2 font-medium text-sm ${
                        activeTab === 'reservations'
                          ? 'border-blue-500 text-blue-600'
                          : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                      }`}
                    >
                      Reservations
                    </button>
                    <button
                      onClick={() => setActiveTab('leases')}
                      className={`py-2 px-1 border-b-2 font-medium text-sm ${
                        activeTab === 'leases'
                          ? 'border-blue-500 text-blue-600'
                          : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                      }`}
                    >
                      Active Leases
                    </button>
                  </nav>
                </div>

                {activeTab === 'reservations' ? (
                  <div className="space-y-3">
                    {reservations.map((reservation) => (
                      <div
                        key={reservation.id}
                        className="p-3 border border-gray-200 rounded-lg"
                      >
                        <div className="flex items-center justify-between">
                          <div>
                            <p className="text-sm font-medium text-gray-900">
                              {reservation.hostname || 'Unnamed'}
                            </p>
                            <p className="text-sm text-gray-600">
                              {reservation.ip_address} → {reservation.mac_address}
                            </p>
                            {reservation.description && (
                              <p className="text-xs text-gray-500 mt-1">
                                {reservation.description}
                              </p>
                            )}
                          </div>
                          <div className="flex items-center space-x-2">
                            <span
                              className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                                reservation.enabled
                                  ? 'bg-green-100 text-green-800'
                                  : 'bg-red-100 text-red-800'
                              }`}
                            >
                              {reservation.enabled ? 'Enabled' : 'Disabled'}
                            </span>
                            <button
                              onClick={() => handleDeleteReservation(reservation.id)}
                              className="text-red-600 hover:text-red-900"
                            >
                              <Trash2 className="h-4 w-4" />
                            </button>
                          </div>
                        </div>
                      </div>
                    ))}
                    {reservations.length === 0 && (
                      <p className="text-gray-500 text-center py-4">
                        No reservations found for this scope
                      </p>
                    )}
                  </div>
                ) : (
                  <div className="space-y-3">
                    {leases.map((lease) => (
                      <div
                        key={lease.id}
                        className="p-3 border border-gray-200 rounded-lg"
                      >
                        <div className="flex items-center justify-between">
                          <div className="flex items-center">
                            <Monitor className="h-5 w-5 text-gray-400 mr-2" />
                            <div>
                              <p className="text-sm font-medium text-gray-900">
                                {lease.hostname || 'Unknown'}
                              </p>
                              <p className="text-sm text-gray-600">
                                {lease.ip_address} → {lease.mac_address}
                              </p>
                              <p className="text-xs text-gray-500">
                                Expires: {new Date(lease.end_time).toLocaleString()}
                              </p>
                            </div>
                          </div>
                          <span
                            className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                              lease.state === 'active'
                                ? 'bg-green-100 text-green-800'
                                : 'bg-yellow-100 text-yellow-800'
                            }`}
                          >
                            {lease.state}
                          </span>
                        </div>
                      </div>
                    ))}
                    {leases.length === 0 && (
                      <p className="text-gray-500 text-center py-4">
                        No active leases found for this scope
                      </p>
                    )}
                  </div>
                )}
              </>
            ) : (
              <p className="text-gray-500 text-center py-8">
                Select a scope to view its details
              </p>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
