import '@/App.scss';
import '@/common/Common.scss';
import MainApp from '@/pages/MainApp';
import GroupListView from '@/pages/GroupListView';
import Login from '@/pages/Login';
import OldApp from '@/_bak/OldApp';
import { useState } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

/**
 * @typedef {Object} UserInfo
 * @property {string} id
 * @property {string} email
 * @property {string} firstName
 * @property {string} lastName
 * @property {string} picture
 * @property {Date} updated_at - Korea Standard Time
 * @property {Date} created_at - Korea Standard Time
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
          <Route path="/" element={<MainApp />} >
            <Route path="group" element={<GroupListView userInfo={userInfo} />} />
          </Route>
          <Route path="/old" element={<OldApp tab="/home"/>} />
        </Routes>
      </BrowserRouter>
    </div>);
}

export default App;
