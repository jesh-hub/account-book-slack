import '@/common/ProcessingSpinner.css';
import {Spinner} from 'react-bootstrap';

function ProcessingSpinner(props) {
  return (
    props.processing ?
      <div className="abs-processing-spinner">
        <Spinner animation="border" variant="secondary" />
      </div> : null
  );
}

export default ProcessingSpinner;
