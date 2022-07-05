import '@/_bak/OldApp.scss';
import Login from '@/pages/Login';
import OldApp from '@/_bak/OldApp';
import { useState } from 'react';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';

/**
 * @typedef {Object} UserInfo
 * @property {string} id
 * @property {string} firstName
 * @property {string} lastName
 * @property {string} picture
 * @property {string} modDate
 * @property {string} regDate
 */

function App() {
  const [userInfo, setUserInfo] = useState();
  if (userInfo === undefined)
    return <Login setUserInfo={setUserInfo} />;
  return (
    <div className="app">
      <BrowserRouter>
        <Routes>
          {/*<Route path="/group" element={<App tab="/ikd"/>} />*/}
          <Route path="/old" element={<OldApp tab="/home"/>} />
          <Route path="*" element={<Navigate replace to="/old" />} />
        </Routes>
      </BrowserRouter>
    </div>);
}

export default App;
