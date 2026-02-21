import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';
import Layout from './components/Layout';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import Templates from './pages/Templates';
import TemplateForm from './pages/TemplateForm';
import ApiKeys from './pages/ApiKeys';
import EmailLogs from './pages/EmailLogs';
import ProtectedRoute from './components/ProtectedRoute';
import PublicRoute from './components/PublicRoute';

export default function App() {
  return (
    <AuthProvider>
        <BrowserRouter>
            <Routes>
                <Route element={<PublicRoute />}>
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                </Route>

                <Route element={<ProtectedRoute />}>
                    <Route element={<Layout />}>
                        <Route path="/" element={<Dashboard />} />
                        <Route path="/templates" element={<Templates />} />
                        <Route path="/templates/new" element={<TemplateForm />} />
                        <Route path="/templates/:id" element={<TemplateForm />} />
                        <Route path="/api-keys" element={<ApiKeys />} />
                        <Route path="/logs" element={<EmailLogs />} />
                    </Route>
                </Route>
            </Routes>
        </BrowserRouter>
    </AuthProvider>
  );
}
