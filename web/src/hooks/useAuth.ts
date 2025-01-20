import { useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';


export const useAuth = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const token = localStorage.getItem('auth_token');

    useEffect(() => {
        if (!token) {
            navigate('/login');
        } else {
            const path = location.pathname === '/login' ? '/' : location.pathname;
            navigate(path);
        } 
    }, [token, location.pathname]);

    return token;
};