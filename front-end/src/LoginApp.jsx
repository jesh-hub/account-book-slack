import '@/LoginApp.scss';
import { useEffect } from 'react';
import ProcessingSpinner from '@/common/ProcessingSpinner';

function LoginApp() {
  async function handleCredentialResponse(response) {
    console.log(response);
  }

  useEffect(() => {
    const js = document.createElement('script');
    js.src = 'https://accounts.google.com/gsi/client';
    js.onerror = () => console.log('error');
    js.onload = () => {
      window.google.accounts.id.initialize({
        client_id: '973841748639-r45e9elajfd2s1h8urlbrqr74kjoll53.apps.googleusercontent.com',
        callback: handleCredentialResponse
      });
      window.google.accounts.id.renderButton(
        document.getElementById('buttonDiv'),
        { theme: 'outline', size: 'large' }
      );
      window.google.accounts.id.prompt();
    };

    document.head.appendChild(js);
  }, []);

  return (
    <div className="login-app">
      <main className="login-app-main">
        <p>
          로그인을 해주세요!
        </p>
        <div id="buttonDiv" />
      </main>
    </div>
  );
}

export default LoginApp;
