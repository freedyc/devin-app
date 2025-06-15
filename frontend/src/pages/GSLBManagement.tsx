import { useEffect, useState } from 'react'
import { Plus, Trash2, Router, Server } from 'lucide-react'
import { gslbApi, type GSLBPolicy, type GSLBPool, type GSLBServer } from '@/services/api'

export default function GSLBManagement() {
  const [policies, setPolicies] = useState<GSLBPolicy[]>([])
  const [selectedPolicy, setSelectedPolicy] = useState<GSLBPolicy | null>(null)
  const [pools, setPools] = useState<GSLBPool[]>([])
  const [selectedPool, setSelectedPool] = useState<GSLBPool | null>(null)
  const [servers, setServers] = useState<GSLBServer[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchPolicies()
  }, [])

  const fetchPolicies = async () => {
    try {
      const response = await gslbApi.getPolicies()
      setPolicies(response.data.policies)
    } catch (error) {
      console.error('Failed to fetch policies:', error)
    } finally {
      setLoading(false)
    }
  }

  const fetchPools = async (policyId: number) => {
    try {
      const response = await gslbApi.getPools(policyId)
      setPools(response.data.pools)
    } catch (error) {
      console.error('Failed to fetch pools:', error)
    }
  }

  const fetchServers = async (poolId: number) => {
    try {
      const response = await gslbApi.getServers(poolId)
      setServers(response.data.servers)
    } catch (error) {
      console.error('Failed to fetch servers:', error)
    }
  }

  const handlePolicySelect = (policy: GSLBPolicy) => {
    setSelectedPolicy(policy)
    setSelectedPool(null)
    setServers([])
    fetchPools(policy.id)
  }

  const handlePoolSelect = (pool: GSLBPool) => {
    setSelectedPool(pool)
    fetchServers(pool.id)
  }

  const handleDeletePolicy = async (id: number) => {
    if (confirm('Are you sure you want to delete this policy?')) {
      try {
        await gslbApi.deletePolicy(id)
        fetchPolicies()
        if (selectedPolicy?.id === id) {
          setSelectedPolicy(null)
          setPools([])
          setSelectedPool(null)
          setServers([])
        }
      } catch (error) {
        console.error('Failed to delete policy:', error)
      }
    }
  }

  const handleDeletePool = async (id: number) => {
    if (confirm('Are you sure you want to delete this pool?')) {
      try {
        await gslbApi.deletePool(id)
        if (selectedPolicy) {
          fetchPools(selectedPolicy.id)
        }
        if (selectedPool?.id === id) {
          setSelectedPool(null)
          setServers([])
        }
      } catch (error) {
        console.error('Failed to delete pool:', error)
      }
    }
  }

  const handleDeleteServer = async (id: number) => {
    if (confirm('Are you sure you want to delete this server?')) {
      try {
        await gslbApi.deleteServer(id)
        if (selectedPool) {
          fetchServers(selectedPool.id)
        }
      } catch (error) {
        console.error('Failed to delete server:', error)
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
        <h1 className="text-2xl font-bold text-gray-900">GSLB Management</h1>
        <p className="mt-1 text-sm text-gray-600">
          Manage Global Server Load Balancing policies, pools, and servers
        </p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                GSLB Policies
              </h3>
              <button className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">
                <Plus className="h-4 w-4 mr-1" />
                Add Policy
              </button>
            </div>

            <div className="space-y-3">
              {policies.map((policy) => (
                <div
                  key={policy.id}
                  className={`p-3 border rounded-lg cursor-pointer transition-colors ${
                    selectedPolicy?.id === policy.id
                      ? 'border-blue-500 bg-blue-50'
                      : 'border-gray-200 hover:border-gray-300'
                  }`}
                  onClick={() => handlePolicySelect(policy)}
                >
                  <div className="flex items-center justify-between">
                    <div className="flex items-center">
                      <Router className="h-5 w-5 text-gray-400 mr-2" />
                      <div>
                        <p className="text-sm font-medium text-gray-900">
                          {policy.name}
                        </p>
                        <p className="text-xs text-gray-500">
                          {policy.fqdn}
                        </p>
                        <p className="text-xs text-gray-500">
                          Method: {policy.method}
                        </p>
                      </div>
                    </div>
                    <div className="flex items-center space-x-2">
                      <span
                        className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                          policy.enabled
                            ? 'bg-green-100 text-green-800'
                            : 'bg-red-100 text-red-800'
                        }`}
                      >
                        {policy.enabled ? 'Enabled' : 'Disabled'}
                      </span>
                      <button
                        onClick={(e) => {
                          e.stopPropagation()
                          handleDeletePolicy(policy.id)
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
                Pools
                {selectedPolicy && (
                  <span className="text-sm font-normal text-gray-500 ml-2">
                    for {selectedPolicy.name}
                  </span>
                )}
              </h3>
              {selectedPolicy && (
                <button className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-green-600 hover:bg-green-700">
                  <Plus className="h-4 w-4 mr-1" />
                  Add Pool
                </button>
              )}
            </div>

            {selectedPolicy ? (
              <div className="space-y-3">
                {pools.map((pool) => (
                  <div
                    key={pool.id}
                    className={`p-3 border rounded-lg cursor-pointer transition-colors ${
                      selectedPool?.id === pool.id
                        ? 'border-green-500 bg-green-50'
                        : 'border-gray-200 hover:border-gray-300'
                    }`}
                    onClick={() => handlePoolSelect(pool)}
                  >
                    <div className="flex items-center justify-between">
                      <div>
                        <p className="text-sm font-medium text-gray-900">
                          {pool.name}
                        </p>
                        <p className="text-xs text-gray-500">
                          Priority: {pool.priority}, Weight: {pool.weight}
                        </p>
                      </div>
                      <div className="flex items-center space-x-2">
                        <span
                          className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                            pool.enabled
                              ? 'bg-green-100 text-green-800'
                              : 'bg-red-100 text-red-800'
                          }`}
                        >
                          {pool.enabled ? 'Enabled' : 'Disabled'}
                        </span>
                        <button
                          onClick={(e) => {
                            e.stopPropagation()
                            handleDeletePool(pool.id)
                          }}
                          className="text-red-600 hover:text-red-900"
                        >
                          <Trash2 className="h-4 w-4" />
                        </button>
                      </div>
                    </div>
                  </div>
                ))}
                {pools.length === 0 && (
                  <p className="text-gray-500 text-center py-4">
                    No pools found for this policy
                  </p>
                )}
              </div>
            ) : (
              <p className="text-gray-500 text-center py-8">
                Select a policy to view its pools
              </p>
            )}
          </div>
        </div>

        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-lg leading-6 font-medium text-gray-900">
                Servers
                {selectedPool && (
                  <span className="text-sm font-normal text-gray-500 ml-2">
                    for {selectedPool.name}
                  </span>
                )}
              </h3>
              {selectedPool && (
                <button className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-purple-600 hover:bg-purple-700">
                  <Plus className="h-4 w-4 mr-1" />
                  Add Server
                </button>
              )}
            </div>

            {selectedPool ? (
              <div className="space-y-3">
                {servers.map((server) => (
                  <div
                    key={server.id}
                    className="p-3 border border-gray-200 rounded-lg"
                  >
                    <div className="flex items-center justify-between">
                      <div className="flex items-center">
                        <Server className="h-5 w-5 text-gray-400 mr-2" />
                        <div>
                          <p className="text-sm font-medium text-gray-900">
                            {server.name}
                          </p>
                          <p className="text-sm text-gray-600">
                            {server.ip_address}:{server.port}
                          </p>
                          <p className="text-xs text-gray-500">
                            Weight: {server.weight}, Priority: {server.priority}
                          </p>
                          <p className="text-xs text-gray-500">
                            Health Check: {server.health_check}
                          </p>
                        </div>
                      </div>
                      <div className="flex items-center space-x-2">
                        <span
                          className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                            server.status === 'healthy'
                              ? 'bg-green-100 text-green-800'
                              : server.status === 'unhealthy'
                              ? 'bg-red-100 text-red-800'
                              : 'bg-yellow-100 text-yellow-800'
                          }`}
                        >
                          {server.status}
                        </span>
                        <span
                          className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                            server.enabled
                              ? 'bg-blue-100 text-blue-800'
                              : 'bg-gray-100 text-gray-800'
                          }`}
                        >
                          {server.enabled ? 'Enabled' : 'Disabled'}
                        </span>
                        <button
                          onClick={() => handleDeleteServer(server.id)}
                          className="text-red-600 hover:text-red-900"
                        >
                          <Trash2 className="h-4 w-4" />
                        </button>
                      </div>
                    </div>
                  </div>
                ))}
                {servers.length === 0 && (
                  <p className="text-gray-500 text-center py-4">
                    No servers found for this pool
                  </p>
                )}
              </div>
            ) : (
              <p className="text-gray-500 text-center py-8">
                Select a pool to view its servers
              </p>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
