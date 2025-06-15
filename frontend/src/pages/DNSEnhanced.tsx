import { useEffect, useState } from 'react'
import { Plus, Trash2, Settings, Upload } from 'lucide-react'
import { dnsEnhancedApi, deploymentApi, type DNSView, type ForwardingZone, type StubZone } from '../services/api'

export default function DNSEnhanced() {
  const [activeTab, setActiveTab] = useState('views')
  const [views, setViews] = useState<DNSView[]>([])
  const [forwardingZones, setForwardingZones] = useState<ForwardingZone[]>([])
  const [stubZones, setStubZones] = useState<StubZone[]>([])
  const [loading, setLoading] = useState(true)
  const [deploying, setDeploying] = useState(false)

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async () => {
    try {
      const [viewsRes, forwardingRes, stubRes] = await Promise.all([
        dnsEnhancedApi.getViews(),
        dnsEnhancedApi.getForwardingZones(),
        dnsEnhancedApi.getStubZones(),
      ])

      setViews(viewsRes.data.views)
      setForwardingZones(forwardingRes.data.zones)
      setStubZones(stubRes.data.zones)
    } catch (error) {
      console.error('Failed to fetch enhanced DNS data:', error)
    } finally {
      setLoading(false)
    }
  }

  const handleDeploy = async () => {
    setDeploying(true)
    try {
      await deploymentApi.deployDNS()
      alert('DNS configuration deployed successfully!')
    } catch (error) {
      console.error('Failed to deploy DNS configuration:', error)
      alert('Failed to deploy DNS configuration')
    } finally {
      setDeploying(false)
    }
  }

  const handleDeleteView = async (id: number) => {
    if (confirm('Are you sure you want to delete this view?')) {
      try {
        await dnsEnhancedApi.deleteView(id)
        fetchData()
      } catch (error) {
        console.error('Failed to delete view:', error)
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
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-bold text-gray-900">Enhanced DNS Management</h1>
            <p className="mt-1 text-sm text-gray-600">
              Manage DNS views, forwarding zones, stub zones, and recursion settings
            </p>
          </div>
          <button
            onClick={handleDeploy}
            disabled={deploying}
            className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-green-600 hover:bg-green-700 disabled:opacity-50"
          >
            <Upload className="h-4 w-4 mr-2" />
            {deploying ? 'Deploying...' : 'Deploy Configuration'}
          </button>
        </div>
      </div>

      <div className="border-b border-gray-200 mb-6">
        <nav className="-mb-px flex space-x-8">
          {[
            { id: 'views', name: 'DNS Views' },
            { id: 'forwarding', name: 'Forwarding Zones' },
            { id: 'stub', name: 'Stub Zones' },
            { id: 'recursion', name: 'Recursion Settings' },
          ].map((tab) => (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id)}
              className={`py-2 px-1 border-b-2 font-medium text-sm ${
                activeTab === tab.id
                  ? 'border-blue-500 text-blue-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
              }`}
            >
              {tab.name}
            </button>
          ))}
        </nav>
      </div>

      {activeTab === 'views' && (
        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                DNS Views
              </h3>
              <button className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">
                <Plus className="h-4 w-4 mr-1" />
                Add View
              </button>
            </div>

            <div className="space-y-3">
              {views.map((view) => (
                <div
                  key={view.id}
                  className="p-3 border border-gray-200 rounded-lg"
                >
                  <div className="flex items-center justify-between">
                    <div>
                      <div className="flex items-center space-x-2">
                        <span className="text-sm font-medium text-gray-900">
                          {view.name}
                        </span>
                        <span className="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-gray-100 text-gray-800">
                          {view.recursion ? 'Recursive' : 'Non-recursive'}
                        </span>
                      </div>
                      <p className="text-sm text-gray-600 mt-1">
                        Match clients: {view.match_clients}
                      </p>
                    </div>
                    <div className="flex items-center space-x-2">
                      <span
                        className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                          view.enabled
                            ? 'bg-green-100 text-green-800'
                            : 'bg-red-100 text-red-800'
                        }`}
                      >
                        {view.enabled ? 'Enabled' : 'Disabled'}
                      </span>
                      <button
                        onClick={() => handleDeleteView(view.id)}
                        className="text-red-600 hover:text-red-900"
                      >
                        <Trash2 className="h-4 w-4" />
                      </button>
                    </div>
                  </div>
                </div>
              ))}
              {views.length === 0 && (
                <p className="text-gray-500 text-center py-4">
                  No DNS views configured
                </p>
              )}
            </div>
          </div>
        </div>
      )}

      {activeTab === 'forwarding' && (
        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                Forwarding Zones
              </h3>
              <button className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">
                <Plus className="h-4 w-4 mr-1" />
                Add Forwarding Zone
              </button>
            </div>

            <div className="space-y-3">
              {forwardingZones.map((zone) => (
                <div
                  key={zone.id}
                  className="p-3 border border-gray-200 rounded-lg"
                >
                  <div className="flex items-center justify-between">
                    <div>
                      <div className="flex items-center space-x-2">
                        <span className="text-sm font-medium text-gray-900">
                          {zone.name}
                        </span>
                        <span className="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-blue-100 text-blue-800">
                          {zone.forward}
                        </span>
                      </div>
                      <p className="text-sm text-gray-600 mt-1">
                        Forwarders: {zone.forwarders.join(', ')}
                      </p>
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
                        onClick={() => dnsEnhancedApi.deleteForwardingZone(zone.id).then(fetchData)}
                        className="text-red-600 hover:text-red-900"
                      >
                        <Trash2 className="h-4 w-4" />
                      </button>
                    </div>
                  </div>
                </div>
              ))}
              {forwardingZones.length === 0 && (
                <p className="text-gray-500 text-center py-4">
                  No forwarding zones configured
                </p>
              )}
            </div>
          </div>
        </div>
      )}

      {activeTab === 'stub' && (
        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                Stub Zones
              </h3>
              <button className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">
                <Plus className="h-4 w-4 mr-1" />
                Add Stub Zone
              </button>
            </div>

            <div className="space-y-3">
              {stubZones.map((zone) => (
                <div
                  key={zone.id}
                  className="p-3 border border-gray-200 rounded-lg"
                >
                  <div className="flex items-center justify-between">
                    <div>
                      <div className="flex items-center space-x-2">
                        <span className="text-sm font-medium text-gray-900">
                          {zone.name}
                        </span>
                        <span className="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-purple-100 text-purple-800">
                          Stub
                        </span>
                      </div>
                      <p className="text-sm text-gray-600 mt-1">
                        Masters: {zone.masters.join(', ')}
                      </p>
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
                        onClick={() => dnsEnhancedApi.deleteStubZone(zone.id).then(fetchData)}
                        className="text-red-600 hover:text-red-900"
                      >
                        <Trash2 className="h-4 w-4" />
                      </button>
                    </div>
                  </div>
                </div>
              ))}
              {stubZones.length === 0 && (
                <p className="text-gray-500 text-center py-4">
                  No stub zones configured
                </p>
              )}
            </div>
          </div>
        </div>
      )}

      {activeTab === 'recursion' && (
        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                Recursion Settings
              </h3>
              <button className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">
                <Settings className="h-4 w-4 mr-1" />
                Configure
              </button>
            </div>

            <div className="space-y-4">
              <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    Recursion Enabled
                  </label>
                  <div className="mt-1">
                    <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                      Yes
                    </span>
                  </div>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700">
                    DNSSEC Validation
                  </label>
                  <div className="mt-1">
                    <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                      Auto
                    </span>
                  </div>
                </div>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">
                  Forwarders
                </label>
                <div className="mt-1">
                  <p className="text-sm text-gray-600">8.8.8.8, 8.8.4.4</p>
                </div>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">
                  Allow Recursion
                </label>
                <div className="mt-1">
                  <p className="text-sm text-gray-600">any</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}
