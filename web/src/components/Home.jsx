
import { NavLink } from 'react-router-dom';

export function Home() {
    
    return (
        <div className="min-h-screen bg-gray-100 items-center justify-center">
            <div className="flex flex-column">
                <div className="text-center">
                    <p className="mt-4 text-gray-600">Authenticated.</p>
                </div>
                <div className="flex flex-row justify-center mt-4 space-x-4">
                    <NavLink to="/profile" className="text-blue-500 hover:underline">Profile</NavLink>
                </div>
            </div>
        </div>
    );
}