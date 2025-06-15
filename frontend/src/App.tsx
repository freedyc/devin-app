import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Layout from './components/Layout'
import Dashboard from './pages/Dashboard'
import DNSManagement from './pages/DNSManagement'
import DNSEnhanced from './pages/DNSEnhanced'
import DHCPManagement from './pages/DHCPManagement'
import GSLBManagement from './pages/GSLBManagement'
import ClusterManagement from './pages/ClusterManagement'
import Configuration from './pages/Configuration'

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/dns" element={<DNSManagement />} />
          <Route path="/dns/enhanced" element={<DNSEnhanced />} />
          <Route path="/dhcp" element={<DHCPManagement />} />
          <Route path="/gslb" element={<GSLBManagement />} />
          <Route path="/cluster" element={<ClusterManagement />} />
          <Route path="/config" element={<Configuration />} />
        </Routes>
      </Layout>
    </Router>
  )
}

export default App
