import { useState, useEffect } from 'react';
import { axiosInstance } from '../utils/axios-instance';
import { LoadingSpinner } from './LoadingSpinner';
import { UserCard } from './UserCard';

export function UserList({ onSelectUser }) {
    const [users, setUsers] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const response = await axiosInstance.get('/allowedUsers');
                setUsers(response.data.results || []);
            } catch (error) {
                console.error('Failed to fetch users:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchUsers();
    }, []);

    return (
        <div className="grid grid-cols-1">
            {loading ? (
                <LoadingSpinner />
            ) : (
                <div className="flex flex-wrap gap-4">
                    {users.map((user) => (
                        <div key={user.id} className="w-80">
                            <UserCard user={user} onSelect={onSelectUser} />
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
}