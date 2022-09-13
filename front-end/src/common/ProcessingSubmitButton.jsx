import { Button, Spinner } from 'react-bootstrap';

export default function ProcessingSubmitButton({
  processing,
  disabled = false,
  children = '확인' }) {
  return <Button
    className="w-100"
    type="submit"
    disabled={processing || disabled}
  >
    {
      processing ?
        <Spinner animation="border" variant="light" size="sm"/> :
        children
    }
  </Button>;
}
