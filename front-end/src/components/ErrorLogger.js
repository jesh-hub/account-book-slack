import { useContext, useState } from 'react';
import { ErrorContext } from '@/common/ErrorContext';
import { Toast } from 'react-bootstrap';

function ErrorToast(props) {
  const [show, setShow] = useState(true);

  function onToastClose() {
    setShow(false);
    props.onToastClose(props.error);
  }

  return (
    <Toast
      show={show}
      onClose={onToastClose}
    >
      <Toast.Header>
        <strong className="me-auto">에러 {props.error.code !== undefined ? `[${props.error.code}]` : ''}</strong>
      </Toast.Header>
      <Toast.Body>{props.error.message}</Toast.Body>
    </Toast>
  );
}

function ErrorLogger() {
  const { errors, deleteError } = useContext(ErrorContext);

  return (
    errors.map(error =>
      <ErrorToast
        key={error.id}
        error={error}
        onToastClose={deleteError}
      />
    ));
}

export default ErrorLogger;
