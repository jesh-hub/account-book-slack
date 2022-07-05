import '@/App.scss';
import Login from '@/pages/Login';
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
          {/*<Route path="/" element={<App tab="home"/>} />*/}
        </Routes>
      </BrowserRouter>
    </div>);
}

export default App;
