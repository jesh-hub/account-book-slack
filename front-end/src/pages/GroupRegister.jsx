import '@/pages/GroupRegister.scss';
import { Badge, Button, Col, Form, Row, Spinner } from 'react-bootstrap';
import { useMemo, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

const BadgeBg = ['info', 'light'];
const BadgeText = [undefined, 'dark'];

function BadgeItem({ index, name, handleRemove }) {
  const bg = useMemo(() => BadgeBg[index % 2], [index]);
  const text = useMemo(() => BadgeText[index % 2], [index]);

  return (
    <>
      <Badge
        size="sm"
        bg={bg}
        text={text}
      >{name}</Badge>
      <Button
        size="xs"
        variant="clear"
        onClick={handleRemove}
      >x</Button>
    </>
  );
}

function RegisteredList({ list, handleRemove }) {
  return (
    <ul>
      {
        list.map((item, i) => (
          <li key={item}>
            <BadgeItem
              index={i}
              name={item}
              handleRemove={() => handleRemove(item)}
            />
          </li>
        ))
      }
    </ul>
  );
}

export default function GroupRegister({ userInfo }) {
  const navigate = useNavigate();
  const [processing, setProcessing] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    paymentMethod: '',
  });
  const [apiData, setApiData] = useState({
    email: [],
    paymentMethod: []
  });

  function handleFormDataChanged(evt) {
    const { name, value } = evt.target;
    setFormData(formData => ({
      ...formData,
      [name]: value
    }));
  }

  function handleApiDataAdded(name) {
    if (apiData[name].includes(formData[name]))
      return;
    setApiData(apiData => ({
      ...apiData,
      [name]: apiData[name].concat(formData[name])
    }));
    setFormData(formData => ({
      ...formData,
      [name]: '',
    }));
  }

  function handleApiDataRemoved(name, value) {
    const index = apiData[name].indexOf(value);
    if (index < 0)
      return;
    setApiData(apiData => ({
      ...apiData,
      [name]: apiData[name].filter(data => data !== value)
    }));
  }

  async function submit(evt) {
    evt.preventDefault();
    try {
      setProcessing(true);
      await axios.post(`${process.env.REACT_APP_ABS}/v1/group`, {
        name: formData.name,
        users: [props.userInfo.email].concat(apiData.email),
        regUserId: props.userInfo.id,
        PaymentMethodAdd: apiData.paymentMethod.map(method => ({
          name: method,
          default: false
        }))
      });
      setProcessing(false);
      navigate(-1);
    } catch (e) {
      setProcessing(false);
      console.log(e);
    }
  }

  return (
    <article className="abs-group-register">
      <Form
        className="register-form"
        onSubmit={submit}
      >
        <Form.Group className="register-row">
          <Form.Control
            name="name"
            type="string"
            value={formData.name}
            placeholder="그룹 이름"
            onChange={handleFormDataChanged}
            required
          />
        </Form.Group>
        {/* 이 밑으로 두 개 form group 비슷한 코드 */}
        <Form.Group className="register-row-group">
          <Form.Label>그룹 구성원 추가하기</Form.Label>
          <Row className="register-row">
            <Col
              as={Form.Control}
              name="email"
              value={formData.email}
              type="email"
              placeholder="초대할 이메일"
              onChange={handleFormDataChanged}
            />
            <Col
              as={Button}
              className="col-3"
              variant="outline-primary"
              // TODO 이메일 형식에 맞지 않으면 disable
              // TODO 이미 추가한 이메일이면 disable
              onClick={() => handleApiDataAdded('email')}
            >
              추가
            </Col>
          </Row>
          <RegisteredList
            list={apiData.email}
            handleRemove={email => handleApiDataRemoved('email', email)}
          />
        </Form.Group>
        <Form.Group className="register-row-group">
          <Form.Label>결제 수단 추가하기</Form.Label>
          <Row className="register-row">
            <Col
              as={Form.Control}
              name="paymentMethod"
              value={formData.paymentMethod}
              type="text"
              placeholder="결제 수단 이름"
              onChange={handleFormDataChanged}
            />
            <Col
              as={Button}
              className="col-3"
              variant="outline-primary"
              // TODO 이미 추가한 결제 수단이면 disable
              onClick={() => handleApiDataAdded('paymentMethod')}
            >
              추가
            </Col>
          </Row>
          <RegisteredList
            list={apiData.paymentMethod}
            handleRemove={method => handleApiDataRemoved('paymentMethod', method)}
          />
        </Form.Group>
        <Button
          className="w-100"
          type="submit"
          disabled={processing}
        >
          {processing ? <Spinner animation="border" variant="light" size="sm" /> : '등록'}
        </Button>
      </Form>
    </article>
  );
}
