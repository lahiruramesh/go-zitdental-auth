import { useState, useEffect } from 'react';
import { axiosInstance } from '../utils/axios-instance';
import { useAuth } from '../hooks/useAuth';


export function Profile() {
  const [profile, setProfile] = useState(null);
  const token = useAuth();

  const fetchProfile = async () => {
    try {

      const response = await axiosInstance.get('/profile', {
          headers: {
              Authorization: `Bearer ${token}`,
          },
      });
      setProfile(response.data.results);
    } catch (error) {
      console.error('Failed to fetch profile:', error);
      //navigate('/login');
      // alert('You are not logged in', error);
    }
  };

  useEffect(() => {
    fetchProfile();
  }, []);

  return (
    <div className="min-h-screen bg-gray-100 p-4">
      <div className="max-w-4xl mx-auto bg-white rounded-lg shadow p-6">
        {profile ? (
          <div className="space-y-4">
            <h1 className="text-2xl font-bold">{profile.username}</h1>
            <div className="grid grid-cols-2 gap-4">
              <div>
                <p className="text-gray-600">Email</p>
                <p>{profile.email}</p>
              </div>
              <div>
                <p className="text-gray-600">Roles</p>
                <p>{profile.roles?.join(', ')}</p>
              </div>
            </div>
          </div>
        ) : (
          <div className="flex justify-center">
            <div className="animate-spin h-8 w-8 border-4 border-blue-500 rounded-full border-t-transparent"></div>
          </div>
        )}
      </div>
    </div>
  );
}