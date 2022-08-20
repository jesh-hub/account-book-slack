import '@/pages/GroupRegisterView.scss';
import { Badge, Button, Col, Form, Row } from 'react-bootstrap';
import { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const BadgeBg = ['info', 'light'];
const BadgeText = [undefined, 'dark'];

export default function GroupRegisterView(props) {
  const navigate = useNavigate();
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

  function handleApiDataChanged(name) {
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

  function removeApiData(name, value) {
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
      await axios.post('/v1/group', {
        name: formData.name,
        users: [props.userInfo.email].concat(apiData.email),
        regUserId: props.userInfo.id,
        modUserId: props.userInfo.id,
        paymentMethod: apiData.paymentMethod
      });
      navigate(-1);
    } catch (e) {
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
        {/* 이 밑으로 두 개 중복 코드 */}
        <Form.Group className="register-row-group">
          <Form.Label>그룹 구성원 추가하기</Form.Label>
          <Row className="register-row">
            <Col
              as={Form.Control}
              name="email"
              value={formData.email}
              placeholder="초대할 이메일"
              onChange={handleFormDataChanged}
            />
            <Col
              as={Button}
              className="col-3"
              variant="outline-primary"
              // TODO 이메일 형식에 맞지 않으면 disable
              // TODO 이미 추가한 이메일이면 disable
              onClick={() => handleApiDataChanged('email')}
            >
              추가
            </Col>
          </Row>
          <ul className="invited-emails">
            {
              apiData.email.map((email, i) => (
                <li
                  key={email}
                  className="li__email"
                >
                  <Badge
                    size="sm"
                    bg={BadgeBg[i % 2]}
                    text={BadgeText[i % 2]}
                  >{email}</Badge>
                  <Button
                    size="xs"
                    variant="clear"
                    onClick={() => removeApiData('email', email)}
                  >x</Button>
                </li>
              ))
            }
          </ul>
        </Form.Group>
        <Form.Group className="register-row-group">
          <Form.Label>결제 수단 추가하기</Form.Label>
          <Row className="register-row">
            <Col
              as={Form.Control}
              name="paymentMethod"
              value={formData.paymentMethod}
              placeholder="결제 수단 이름"
              onChange={handleFormDataChanged}
            />
            <Col
              as={Button}
              className="col-3"
              variant="outline-primary"
              // TODO 이미 추가한 결제 수단이면 disable
              onClick={() => handleApiDataChanged('paymentMethod')}
            >
              추가
            </Col>
          </Row>
          <ul className="invited-emails">
            {
              apiData.paymentMethod.map((method, i) => (
                <li
                  key={method}
                  className="li__email"
                >
                  <Badge
                    size="sm"
                    bg={BadgeBg[i % 2]}
                    text={BadgeText[i % 2]}
                  >{method}</Badge>
                  <Button
                    size="xs"
                    variant="clear"
                    onClick={() => removeApiData('paymentMethod', method)}
                  >x</Button>
                </li>
              ))
            }
          </ul>
        </Form.Group>
        <Button
          className="w-100"
          type="submit"
        >
          등록
        </Button>
      </Form>
    </article>
  );
}
