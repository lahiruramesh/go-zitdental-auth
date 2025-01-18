import { useState, useEffect } from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import { UserList } from './components/UserList';
import { UserLogin } from './components/UserLogin';
import { AuthCallback } from './components/AuthCallback';
import { Profile } from './components/Profile';
import { Home } from './components/Home';

function App() {
  const [selectedUser, setSelectedUser] = useState(null);
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('auth_token');
    if (token) {
      setIsAuthenticated(true);
    } else {
      setIsAuthenticated(false);
    }
  }, []);

  const handleLogin = (token) => {
    // localStorage.setItem('auth_token', token);
    setIsAuthenticated(true);
  }

  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/callback" element={<AuthCallback />} />
      <Route path="/login" element={
        <div className="min-h-screen bg-gray-100 flex items-center justify-center p-4">
          <div className="grid md:grid-cols-2 gap-4 items-start">
            <div className="w-full">
              <UserList onSelectUser={setSelectedUser} />
            </div>
            {selectedUser && (
              <div className="w-full">
                <UserLogin
                  user={selectedUser}
                  onLogin={handleLogin}
                  onBack={() => setSelectedUser(null)}
                />
              </div>
            )}
          </div>
        </div>
      } />
      <Route 
        path="/profile" 
        element={isAuthenticated ? <Profile /> : <Navigate to="/login" />} 
      />
    </Routes>
  );
}

export default App;