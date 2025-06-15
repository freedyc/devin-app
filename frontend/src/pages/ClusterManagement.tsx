import { useEffect, useState } from 'react'
import { Plus, Trash2, Server, RefreshCw } from 'lucide-react'
import { clusterApi, type ClusterNode } from '../services/api'

export default function ClusterManagement() {
  const [nodes, setNodes] = useState<ClusterNode[]>([])
  const [loading, setLoading] = useState(true)
  const [syncing, setSyncing] = useState(false)

  useEffect(() => {
    fetchNodes()
  }, [])

  const fetchNodes = async () => {
    try {
      const response = await clusterApi.getNodes()
      setNodes(response.data.nodes)
    } catch (error) {
      console.error('Failed to fetch cluster nodes:', error)
    } finally {
      setLoading(false)
    }
  }

  const handleSync = async () => {
    setSyncing(true)
    try {
      await clusterApi.syncConfiguration()
      alert('Cluster configuration synchronized successfully!')
    } catch (error) {
      console.error('Failed to sync cluster configuration:', error)
      alert('Failed to sync cluster configuration')
    } finally {
      setSyncing(false)
    }
  }

  const handleDeleteNode = async (id: number) => {
    if (confirm('Are you sure you want to delete this node?')) {
      try {
        await clusterApi.deleteNode(id)
        fetchNodes()
      } catch (error) {
        console.error('Failed to delete node:', error)
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
            <h1 className="text-2xl font-bold text-gray-900">Cluster Management</h1>
            <p className="mt-1 text-sm text-gray-600">
              Manage cluster nodes and synchronize configurations
            </p>
          </div>
          <div className="flex space-x-3">
            <button
              onClick={handleSync}
              disabled={syncing}
              className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-green-600 hover:bg-green-700 disabled:opacity-50"
            >
              <RefreshCw className="h-4 w-4 mr-2" />
              {syncing ? 'Syncing...' : 'Sync Configuration'}
            </button>
            <button className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">
              <Plus className="h-4 w-4 mr-2" />
              Add Node
            </button>
          </div>
        </div>
      </div>

      <div className="bg-white shadow rounded-lg">
        <div className="px-4 py-5 sm:p-6">
          <div className="space-y-3">
            {nodes.map((node) => (
              <div
                key={node.id}
                className="p-4 border border-gray-200 rounded-lg"
              >
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-3">
                    <Server className="h-8 w-8 text-gray-400" />
                    <div>
                      <div className="flex items-center space-x-2">
                        <span className="text-lg font-medium text-gray-900">
                          {node.name}
                        </span>
                        <span className={`inline-flex items-center px-2 py-0.5 rounded text-xs font-medium ${
                          node.role === 'primary' 
                            ? 'bg-blue-100 text-blue-800' 
                            : 'bg-gray-100 text-gray-800'
                        }`}>
                          {node.role}
                        </span>
                      </div>
                      <p className="text-sm text-gray-600">
                        IP: {node.ip_address}
                      </p>
                    </div>
                  </div>
                  <div className="flex items-center space-x-3">
                    <span
                      className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                        node.status === 'active'
                          ? 'bg-green-100 text-green-800'
                          : 'bg-red-100 text-red-800'
                      }`}
                    >
                      {node.status}
                    </span>
                    {node.role !== 'primary' && (
                      <button
                        onClick={() => handleDeleteNode(node.id)}
                        className="text-red-600 hover:text-red-900"
                      >
                        <Trash2 className="h-4 w-4" />
                      </button>
                    )}
                  </div>
                </div>
              </div>
            ))}
            {nodes.length === 0 && (
              <p className="text-gray-500 text-center py-8">
                No cluster nodes configured
              </p>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
