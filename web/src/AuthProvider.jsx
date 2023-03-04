import React from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

const AuthContext = React.createContext();

const parseHash = (hash) => {
  const data = hash.substring(1)

  if (!data) {
    return null;
  }

  return data.split('&').reduce((parsed, pair) => {
    const params = { ...parsed };
    const [key, value] = pair.split('=');
    params[decodeURIComponent(key)] = decodeURIComponent(value);
    return params;
  }, {});
}

export const AuthProvider = ({ children }) => {
  const [auth, setAuth] = React.useState(null);
  const location = useLocation();
  const navigate = useNavigate();

  React.useEffect(() => {
    const data = parseHash(location.hash);
    if (data === null) return;

    setAuth(data);
    navigate('/');

  }, [auth, location.hash, navigate]);


  return <AuthContext.Provider value={{ auth, setAuth }}>{children}</AuthContext.Provider>;
};

const useAuth = () => React.useContext(AuthContext);

export default useAuth;
