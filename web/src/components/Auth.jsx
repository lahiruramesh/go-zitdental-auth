import React, { useState } from 'react';
import { UserList } from './UserList';
import { UserLogin } from './UserLogin';

export function Auth() {
    const [selectedUser, setSelectedUser] = useState(null);
    
    return (
        <div className="min-h-screen bg-gray-100 flex items-center justify-center p-4">
            <div className="grid md:grid-cols-2 gap-4 items-start">
                <div className="w-full">
                    <UserList onSelectUser={setSelectedUser} />
                </div>
                {selectedUser && (
                    <div className="w-full">
                        <UserLogin
                            user={selectedUser}
                        />
                    </div>
                )}
            </div>
        </div>
    )
}