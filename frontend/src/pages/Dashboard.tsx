import { useEffect, useState } from 'react'
import { Globe, Network, Router, Activity } from 'lucide-react'
import { dnsApi, dhcpApi, gslbApi, type DNSZone, type DHCPScope, type GSLBPolicy } from '@/services/api'

export default function Dashboard() {
  const [stats, setStats] = useState({
    zones: 0,
    scopes: 0,
    policies: 0,
    activeLeases: 0,
  })
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const [zonesRes, scopesRes, policiesRes] = await Promise.all([
          dnsApi.getZones(),
          dhcpApi.getScopes(),
          gslbApi.getPolicies(),
        ])

        setStats({
          zones: zonesRes.data.total,
          scopes: scopesRes.data.total,
          policies: policiesRes.data.total,
          activeLeases: 0,
        })
      } catch (error) {
        console.error('Failed to fetch dashboard stats:', error)
      } finally {
        setLoading(false)
      }
    }

    fetchStats()
  }, [])

  const statCards = [
    {
      title: 'DNS Zones',
      value: stats.zones,
      icon: Globe,
      color: 'bg-blue-500',
    },
    {
      title: 'DHCP Scopes',
      value: stats.scopes,
      icon: Network,
      color: 'bg-green-500',
    },
    {
      title: 'GSLB Policies',
      value: stats.policies,
      icon: Router,
      color: 'bg-purple-500',
    },
    {
      title: 'Active Leases',
      value: stats.activeLeases,
      icon: Activity,
      color: 'bg-orange-500',
    },
  ]

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
        <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>
        <p className="mt-1 text-sm text-gray-600">
          Overview of your DNS, DHCP, and GSLB infrastructure
        </p>
      </div>

      <div className="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
        {statCards.map((card) => {
          const Icon = card.icon
          return (
            <div
              key={card.title}
              className="relative bg-white pt-5 px-4 pb-12 sm:pt-6 sm:px-6 shadow rounded-lg overflow-hidden"
            >
              <dt>
                <div className={`absolute ${card.color} rounded-md p-3`}>
                  <Icon className="h-6 w-6 text-white" />
                </div>
                <p className="ml-16 text-sm font-medium text-gray-500 truncate">
                  {card.title}
                </p>
              </dt>
              <dd className="ml-16 pb-6 flex items-baseline sm:pb-7">
                <p className="text-2xl font-semibold text-gray-900">
                  {card.value}
                </p>
              </dd>
            </div>
          )
        })}
      </div>

      <div className="mt-8 grid grid-cols-1 gap-5 lg:grid-cols-2">
        <div className="bg-white overflow-hidden shadow rounded-lg">
          <div className="p-5">
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <Globe className="h-6 w-6 text-gray-400" />
              </div>
              <div className="ml-5 w-0 flex-1">
                <dl>
                  <dt className="text-sm font-medium text-gray-500 truncate">
                    DNS Status
                  </dt>
                  <dd className="text-lg font-medium text-gray-900">
                    All zones operational
                  </dd>
                </dl>
              </div>
            </div>
          </div>
          <div className="bg-gray-50 px-5 py-3">
            <div className="text-sm">
              <span className="font-medium text-green-600">Healthy</span>
              <span className="text-gray-500"> - All DNS services running</span>
            </div>
          </div>
        </div>

        <div className="bg-white overflow-hidden shadow rounded-lg">
          <div className="p-5">
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <Network className="h-6 w-6 text-gray-400" />
              </div>
              <div className="ml-5 w-0 flex-1">
                <dl>
                  <dt className="text-sm font-medium text-gray-500 truncate">
                    DHCP Status
                  </dt>
                  <dd className="text-lg font-medium text-gray-900">
                    All scopes active
                  </dd>
                </dl>
              </div>
            </div>
          </div>
          <div className="bg-gray-50 px-5 py-3">
            <div className="text-sm">
              <span className="font-medium text-green-600">Healthy</span>
              <span className="text-gray-500"> - DHCP server operational</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
