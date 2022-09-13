import '@/App.scss';
import '@/common/Common.scss';
import MainApp from '@/pages/MainApp';
import GroupList from '@/pages/GroupList';
import GroupRegister from '@/pages/GroupRegister';
import Login from '@/pages/Login';
import PaymentList from '@/pages/PaymentList';
import PaymentRegister from '@/pages/PaymentRegister';
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
            <Route path="groups" element={<GroupList userInfo={userInfo} />} />
            <Route path="groups/register" element={<GroupRegister userInfo={userInfo} />} />
            <Route path="payments" element={<PaymentList />} />
            <Route path="payments/register" element={<PaymentRegister userInfo={userInfo} />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </div>);
}

export default App;
