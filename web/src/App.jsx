import { Route, Routes } from 'react-router-dom';

import { Auth } from './components/Auth';
import { AuthCallback } from './components/AuthCallback';
import { Profile } from './components/Profile';
import { Home } from './components/Home';

import { useAuth } from './hooks/useAuth';

function App() {
  const token = useAuth();

  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/callback" element={<AuthCallback />} />
      <Route path="/login" element={<Auth />} />
      <Route path="/profile" element={<Profile />} />
    </Routes>
  );
}

export default App;