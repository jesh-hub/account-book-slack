import '@/LoginApp.scss';
import { useEffect, useState } from 'react';
import axios from 'axios';
import ProcessingSpinner from '@/common/ProcessingSpinner';

const GoogleSrc = 'https://accounts.google.com/gsi/client';

function LoginApp() {
  const [processing, setProcessing] = useState(false);

  async function handleCredentialResponse({ credential }) {
    setProcessing(true);
    try {
      const res = await axios.post(`${process.env.REACT_APP_ABS}/v1/login`, { credential });
      console.log(res);
    } catch (e) {
      console.error(e);
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
        client_id: '973841748639-r45e9elajfd2s1h8urlbrqr74kjoll53.apps.googleusercontent.com',
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
  }, []);

  return (
    <div className="login-app">
      <main className="app-main">
        <p>
          <b>로그인</b>을 해주세요!
        </p>
        <div id="buttonDiv" />
        <ProcessingSpinner processing={processing} />
      </main>
    </div>
  );
}

export default LoginApp;
