import '@/pages/GroupRegister.scss';
import { Badge, Button, Col, Form, Row } from 'react-bootstrap';
import { useCallback, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { doPostRequest } from '@/common/Api';
import ProcessingSubmitButton from '@/components/ProcessingSubmitButton';

const BadgeBg = ['info', 'light'];
const BadgeText = [undefined, 'dark'];

function RegisteredList({ list, handleRemove }) {
  return (
    <ul>
      {
        list.map((item, i) => (
          <li key={item}>
            <Badge
              size="sm"
              bg={BadgeBg[i % 2]}
              text={BadgeText[i % 2]}
            >{item}</Badge>
            <Button
              size="xs"
              variant="clear"
              onClick={() => handleRemove(item)}
            >x</Button>
          </li>
        ))
      }
    </ul>
  );
}

function RegisterFormRow({ input, handleFormDataChanged, handleApiDataAdded }) {
  const { name, type, value, placeholder } = input;

  return (
    <Row className="register-row">
      <Col
        as={Form.Control}
        name={name}
        value={value}
        type={type}
        placeholder={placeholder}
        onChange={handleFormDataChanged}
      />
      <Col
        as={Button}
        className="col-3"
        variant="outline-primary"
        // TODO 이 외의 disable 조건 적용
        disabled={value.length === 0}
        onClick={() => handleApiDataAdded(name)}
      >
        추가
      </Col>
    </Row>
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

  const handleFormDataChanged = useCallback(evt => {
    const { name, value } = evt.target;
    setFormData(formData => ({
      ...formData,
      [name]: value
    }));
  }, []);

  const handleApiDataAdded = useCallback(name => {
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
  }, [apiData, formData]);

  const handleApiDataRemoved = useCallback((name, value) => {
    const index = apiData[name].indexOf(value);
    if (index < 0)
      return;
    setApiData(apiData => ({
      ...apiData,
      [name]: apiData[name].filter(data => data !== value)
    }));
  }, [apiData]);

  async function submit(evt) {
    evt.preventDefault();
    try {
      setProcessing(true);
      await doPostRequest('/v1/group', {
        name: formData.name,
        users: [userInfo.email].concat(apiData.email),
        regUserId: userInfo.id,
        paymentMethodAdd: apiData.paymentMethod.map((method, i) => ({
          name: method,
          default: i === 0
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
            type="text"
            value={formData.name}
            placeholder="그룹 이름"
            onChange={handleFormDataChanged}
            required
          />
        </Form.Group>
        <Form.Group className="register-row-group">
          <Form.Label>그룹 구성원 추가하기</Form.Label>
          <RegisterFormRow
            input={{
              name: 'email',
              type: 'email',
              value: formData.email,
              placeholder: '초대할 구성원 이메일',
            }}
            handleFormDataChanged={handleFormDataChanged}
            handleApiDataAdded={handleApiDataAdded}
          />
          <RegisteredList
            list={apiData.email}
            handleRemove={email => handleApiDataRemoved('email', email)}
          />
        </Form.Group>
        <Form.Group className="register-row-group">
          <Form.Label>결제 수단 추가하기</Form.Label>
          <RegisterFormRow
            input={{
              name: 'paymentMethod',
              type: 'text',
              value: formData.paymentMethod,
              placeholder: '결제 수단 이름'
            }}
            handleFormDataChanged={handleFormDataChanged}
            handleApiDataAdded={handleApiDataAdded}
          />
          <RegisteredList
            list={apiData.paymentMethod}
            handleRemove={method => handleApiDataRemoved('paymentMethod', method)}
          />
        </Form.Group>
        <ProcessingSubmitButton processing={processing}>등록</ProcessingSubmitButton>
      </Form>
    </article>
  );
}
