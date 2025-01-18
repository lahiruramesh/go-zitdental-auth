import { NavLink } from 'react-router-dom';
export function Home() {

    return (
        <div className="min-h-screen bg-gray-100 flex flex-row items-center justify-center">
            <div className="text-center">
                <p className="mt-4 text-gray-600">Authenticated.</p>
            </div>
            <div className="flex justify-center mt-4 space-x-4">
                <NavLink to="/profile" className="text-blue-500 hover:underline">Profile</NavLink>
                <NavLink to="/login" className="text-blue-500 hover:underline">Login</NavLink>
            </div>
        </div>
    );
}