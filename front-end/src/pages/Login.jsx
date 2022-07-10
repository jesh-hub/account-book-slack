import '@/pages/Login.scss';
import axios from 'axios';
import ErrorContext from '@/common/ErrorContext';
import ProcessingSpinner from '@/common/ProcessingSpinner';
import { useContext, useEffect, useState } from 'react';
import ErrorLogger from '@/components/ErrorLogger';

const GoogleSrc = 'https://accounts.google.com/gsi/client';

export default function Login(props) {
  const [processing, setProcessing] = useState(false);
  const { addError } = useContext(ErrorContext);

  async function handleCredentialResponse({ credential }) {
    setProcessing(true);
    try {
      const { data } = await axios.post(`${process.env.REACT_APP_ABS}/v1/login`, { credential });
      const userInfo = {
        id: data.id,
        firstName: data.firstName,
        lastName: data.lastName,
        picture: data.picture,
        modDate: new Date(data.modDate),
        regDate: new Date(data.regDate)
      };
      window.localStorage.setItem('ABS_userInfo', JSON.stringify(userInfo));
      props.setUserInfo(userInfo);
    } catch (e) {
      addError(e);
    } finally {
      setProcessing(false);
    }
  }

  useEffect(() => {
    const js = document.createElement('script');
    js.src = GoogleSrc;
    js.onerror = () => console.log('error');
    js.onload = () => {
      window.google.accounts.id.initialize({
        client_id: process.env.GOOGLE_CLIENT_ID,
        callback: handleCredentialResponse
      });
      window.google.accounts.id.renderButton(
        document.getElementById('buttonDiv'),
        {
          theme: 'outline',
          shape: 'pill',
          type: 'standard',
          text: 'continue_with',
          size: 'large'
        }
      );
      window.google.accounts.id.prompt();
      document.head.removeChild(js);
    };
    document.head.appendChild(js);
  // eslint-disable-next-line
  }, []);

  return (
    <>
      <section className="login-section">
        <h4 className="section-title">
          <b>로그인</b>을 해주세요!
        </h4>
        <div id="buttonDiv" />
        <ProcessingSpinner processing={processing} />
      </section>
      <aside className="app-aside">
        <ErrorLogger />
      </aside>
    </>
  );
}
