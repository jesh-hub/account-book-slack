import '@/App.scss';
import '@/common/Common.scss';
import MainApp from '@/pages/MainApp';
import GroupListView from '@/pages/GroupListView';
import GroupRegisterView from '@/pages/GroupRegisterView';
import Login from '@/pages/Login';
import PaymentListView from '@/pages/PaymentListView';
import PaymentRegisterView from '@/pages/PaymentRegisterView';
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
            <Route path="groups" element={<GroupListView userInfo={userInfo} />} />
            <Route path="groups/register" element={<GroupRegisterView userInfo={userInfo} />} />
            <Route path="payments" element={<PaymentListView />} />
            <Route path="payments/register" element={<PaymentRegisterView userInfo={userInfo} />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </div>);
}

export default App;
