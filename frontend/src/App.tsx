import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Layout from './components/Layout'
import Dashboard from './pages/Dashboard'
import DNSManagement from './pages/DNSManagement'
import DHCPManagement from './pages/DHCPManagement'
import GSLBManagement from './pages/GSLBManagement'
import Configuration from './pages/Configuration'

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/dns" element={<DNSManagement />} />
          <Route path="/dhcp" element={<DHCPManagement />} />
          <Route path="/gslb" element={<GSLBManagement />} />
          <Route path="/config" element={<Configuration />} />
        </Routes>
      </Layout>
    </Router>
  )
}

export default App
