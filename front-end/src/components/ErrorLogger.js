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

function ErrorLogger(props) {
  // const { errors, deleteError } = useContext(ErrorContext);

  return (
    props.errors.map(error =>
      <ErrorToast
        key={error.id}
        error={error}
        onToastClose={props.deleteError}
      />
    ));
}

export default ErrorLogger;
