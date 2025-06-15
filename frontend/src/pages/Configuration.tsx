import { useState } from 'react'
import { RefreshCw, CheckCircle, XCircle, Settings, AlertTriangle } from 'lucide-react'
import { configApi } from '@/services/api'

export default function Configuration() {
  const [loading, setLoading] = useState({
    bind9Reload: false,
    dhcpRestart: false,
    bind9Validate: false,
    dhcpValidate: false,
  })
  const [results, setResults] = useState<{
    [key: string]: { success: boolean; message: string; output?: string }
  }>({})

  const handleBind9Reload = async () => {
    setLoading(prev => ({ ...prev, bind9Reload: true }))
    try {
      const response = await configApi.reloadBind9()
      setResults(prev => ({
        ...prev,
        bind9Reload: {
          success: true,
          message: response.data.message,
          output: response.data.output,
        },
      }))
    } catch (error: any) {
      setResults(prev => ({
        ...prev,
        bind9Reload: {
          success: false,
          message: error.response?.data?.error || 'Failed to reload BIND9',
          output: error.response?.data?.output,
        },
      }))
    } finally {
      setLoading(prev => ({ ...prev, bind9Reload: false }))
    }
  }

  const handleDHCPRestart = async () => {
    setLoading(prev => ({ ...prev, dhcpRestart: true }))
    try {
      const response = await configApi.restartDHCP()
      setResults(prev => ({
        ...prev,
        dhcpRestart: {
          success: true,
          message: response.data.message,
          output: response.data.output,
        },
      }))
    } catch (error: any) {
      setResults(prev => ({
        ...prev,
        dhcpRestart: {
          success: false,
          message: error.response?.data?.error || 'Failed to restart DHCP',
          output: error.response?.data?.output,
        },
      }))
    } finally {
      setLoading(prev => ({ ...prev, dhcpRestart: false }))
    }
  }

  const handleBind9Validate = async () => {
    setLoading(prev => ({ ...prev, bind9Validate: true }))
    try {
      const response = await configApi.validateBind9()
      setResults(prev => ({
        ...prev,
        bind9Validate: {
          success: response.data.valid,
          message: response.data.message,
          output: response.data.output,
        },
      }))
    } catch (error: any) {
      setResults(prev => ({
        ...prev,
        bind9Validate: {
          success: false,
          message: error.response?.data?.error || 'Failed to validate BIND9 config',
          output: error.response?.data?.output,
        },
      }))
    } finally {
      setLoading(prev => ({ ...prev, bind9Validate: false }))
    }
  }

  const handleDHCPValidate = async () => {
    setLoading(prev => ({ ...prev, dhcpValidate: true }))
    try {
      const response = await configApi.validateDHCP()
      setResults(prev => ({
        ...prev,
        dhcpValidate: {
          success: response.data.valid,
          message: response.data.message,
          output: response.data.output,
        },
      }))
    } catch (error: any) {
      setResults(prev => ({
        ...prev,
        dhcpValidate: {
          success: false,
          message: error.response?.data?.error || 'Failed to validate DHCP config',
          output: error.response?.data?.output,
        },
      }))
    } finally {
      setLoading(prev => ({ ...prev, dhcpValidate: false }))
    }
  }

  const ActionButton = ({ 
    onClick, 
    loading, 
    children, 
    variant = 'primary' 
  }: { 
    onClick: () => void
    loading: boolean
    children: React.ReactNode
    variant?: 'primary' | 'secondary' | 'danger'
  }) => {
    const baseClasses = "inline-flex items-center px-4 py-2 border text-sm font-medium rounded-md focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:opacity-50"
    const variantClasses = {
      primary: "border-transparent text-white bg-blue-600 hover:bg-blue-700 focus:ring-blue-500",
      secondary: "border-gray-300 text-gray-700 bg-white hover:bg-gray-50 focus:ring-blue-500",
      danger: "border-transparent text-white bg-red-600 hover:bg-red-700 focus:ring-red-500"
    }

    return (
      <button
        onClick={onClick}
        disabled={loading}
        className={`${baseClasses} ${variantClasses[variant]}`}
      >
        {loading ? (
          <RefreshCw className="h-4 w-4 mr-2 animate-spin" />
        ) : (
          <Settings className="h-4 w-4 mr-2" />
        )}
        {children}
      </button>
    )
  }

  const ResultDisplay = ({ result }: { result?: { success: boolean; message: string; output?: string } }) => {
    if (!result) return null

    return (
      <div className={`mt-3 p-3 rounded-md ${result.success ? 'bg-green-50 border border-green-200' : 'bg-red-50 border border-red-200'}`}>
        <div className="flex items-center">
          {result.success ? (
            <CheckCircle className="h-5 w-5 text-green-400 mr-2" />
          ) : (
            <XCircle className="h-5 w-5 text-red-400 mr-2" />
          )}
          <p className={`text-sm font-medium ${result.success ? 'text-green-800' : 'text-red-800'}`}>
            {result.message}
          </p>
        </div>
        {result.output && (
          <div className="mt-2">
            <pre className={`text-xs p-2 rounded ${result.success ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'} overflow-x-auto`}>
              {result.output}
            </pre>
          </div>
        )}
      </div>
    )
  }

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-gray-900">Configuration</h1>
        <p className="mt-1 text-sm text-gray-600">
          Manage and validate DNS and DHCP service configurations
        </p>
      </div>

      <div className="space-y-6">
        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <h3 className="text-lg leading-6 font-medium text-gray-900 mb-4">
              BIND9 DNS Configuration
            </h3>
            
            <div className="space-y-4">
              <div className="flex items-center justify-between p-4 border border-gray-200 rounded-lg">
                <div>
                  <h4 className="text-sm font-medium text-gray-900">Validate Configuration</h4>
                  <p className="text-sm text-gray-500">Check BIND9 configuration files for syntax errors</p>
                </div>
                <ActionButton
                  onClick={handleBind9Validate}
                  loading={loading.bind9Validate}
                  variant="secondary"
                >
                  Validate
                </ActionButton>
              </div>
              <ResultDisplay result={results.bind9Validate} />

              <div className="flex items-center justify-between p-4 border border-gray-200 rounded-lg">
                <div>
                  <h4 className="text-sm font-medium text-gray-900">Reload Configuration</h4>
                  <p className="text-sm text-gray-500">Reload BIND9 configuration without restarting the service</p>
                </div>
                <ActionButton
                  onClick={handleBind9Reload}
                  loading={loading.bind9Reload}
                >
                  Reload
                </ActionButton>
              </div>
              <ResultDisplay result={results.bind9Reload} />
            </div>
          </div>
        </div>

        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <h3 className="text-lg leading-6 font-medium text-gray-900 mb-4">
              DHCP Configuration
            </h3>
            
            <div className="space-y-4">
              <div className="flex items-center justify-between p-4 border border-gray-200 rounded-lg">
                <div>
                  <h4 className="text-sm font-medium text-gray-900">Validate Configuration</h4>
                  <p className="text-sm text-gray-500">Check DHCP configuration files for syntax errors</p>
                </div>
                <ActionButton
                  onClick={handleDHCPValidate}
                  loading={loading.dhcpValidate}
                  variant="secondary"
                >
                  Validate
                </ActionButton>
              </div>
              <ResultDisplay result={results.dhcpValidate} />

              <div className="flex items-center justify-between p-4 border border-yellow-200 rounded-lg bg-yellow-50">
                <div className="flex items-start">
                  <AlertTriangle className="h-5 w-5 text-yellow-400 mr-2 mt-0.5" />
                  <div>
                    <h4 className="text-sm font-medium text-gray-900">Restart DHCP Service</h4>
                    <p className="text-sm text-gray-500">Restart the DHCP service to apply configuration changes</p>
                    <p className="text-xs text-yellow-600 mt-1">Warning: This will temporarily interrupt DHCP service</p>
                  </div>
                </div>
                <ActionButton
                  onClick={handleDHCPRestart}
                  loading={loading.dhcpRestart}
                  variant="danger"
                >
                  Restart
                </ActionButton>
              </div>
              <ResultDisplay result={results.dhcpRestart} />
            </div>
          </div>
        </div>

        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <h3 className="text-lg leading-6 font-medium text-gray-900 mb-4">
              System Information
            </h3>
            
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div className="p-4 border border-gray-200 rounded-lg">
                <h4 className="text-sm font-medium text-gray-900">BIND9 Status</h4>
                <p className="text-sm text-gray-500 mt-1">DNS service is running</p>
                <div className="mt-2">
                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                    Active
                  </span>
                </div>
              </div>

              <div className="p-4 border border-gray-200 rounded-lg">
                <h4 className="text-sm font-medium text-gray-900">DHCP Status</h4>
                <p className="text-sm text-gray-500 mt-1">DHCP service is running</p>
                <div className="mt-2">
                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                    Active
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
