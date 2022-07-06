import '@/App.scss';
import Login from '@/pages/Login';
import OldApp from '@/_bak/OldApp';
import { useState } from 'react';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';

/**
 * @typedef {Object} UserInfo
 * @property {string} id - email 형식
 * @property {string} firstName
 * @property {string} lastName
 * @property {string} picture
 * @property {Date} modDate - Korea Standard Time
 * @property {Date} regDate - Korea Standard Time
 */

function App() {
  const _userInfo = window.localStorage.getItem('ABS_userInfo');
  const [userInfo, setUserInfo] = useState(JSON.parse(_userInfo) || undefined);

  if (userInfo === undefined)
    return <Login setUserInfo={setUserInfo} />;
  return (
    <div className="abs-app">
      <BrowserRouter>
        <Routes>
          <Route path="/old" element={<OldApp tab="/home"/>} />
          <Route path="*" element={<Navigate replace to="/old" />} />
        </Routes>
      </BrowserRouter>
    </div>);
}

export default App;
