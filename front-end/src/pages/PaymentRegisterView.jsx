import '@/pages/PaymentRegisterView.scss';
import { Button, ButtonGroup, Col, Dropdown, DropdownButton, Form } from 'react-bootstrap';
import { useState } from 'react';
import { useLocation } from 'react-router-dom';
import axios from 'axios';
import useRequest from '@/common/useRequest';

const BtnTypes = [
  { key: 'income', uiText: '지출', value: -1 },
  { key: 'outgoing', uiText: '수입', value: 1 },
];

function TypeButtons(props) {
  return BtnTypes.map(btn =>
    <Button
      variant="outline-primary"
      size="sm"
      key={btn.key}
      active={props.active === btn}
      onClick={() => props.setActiveBtn(btn)}
    >
      {btn.uiText}
    </Button>);
}

function DropdownPaymentMethods(props) {
  const [isWaitingPaymentMethods, paymentMethods] = useRequest(
    '/v1/paymentMethod', { groupId: props.gid }, [], []);
  return (
    <DropdownButton
      title={props.active?.name || '(선택)'}
      disabled={isWaitingPaymentMethods}
      size="sm"
      variant="outline-primary"
    >
      {
        paymentMethods.map((method, i) =>
          <Dropdown.Item
            key={i}
            onClick={() => props.setActiveMethod(method)}
          >
            {method.name}
          </Dropdown.Item>)
      }
    </DropdownButton>
  );
}

export default function PaymentRegisterView(props) {
  const location = useLocation();

  const now = new Date();
  const [nowStr, setNowStr] = useState(
    `${now.getFullYear()}-${_zeroPad(now.getMonth() + 1)}-${_zeroPad(now.getDate())}` +
    `T${_zeroPad(now.getHours())}:${_zeroPad(now.getMinutes())}`);
  const [name, setName] = useState('');
  const [activeBtn, setActiveBtn] = useState(BtnTypes[0]);
  const [price, setPrice] = useState('');
  const [activeMethod, setActiveMethod] = useState();
  const [monthlyInstallment, setMonthlyInstallment] = useState('');
  const [category, setCategory] = useState('');

  async function submit(arg) {
    arg.preventDefault();
    try {
      await axios.post(`${process.env.REACT_APP_ABS}/v1/payment`, {
        date: nowStr + ':00+09:00',
        name,
        category,
        price: price * activeBtn.value,
        monthlyInstallment: monthlyInstallment || 0,
        paymentMethodId: activeMethod.id,
        groupId: location.state.gid,
        regUserId: props.userInfo.id
      });
    } catch (e) {
      console.log(e);
    }
  }

  return (
    <article className="abs-register">
      <Form onSubmit={submit}>
        <Form.Group className="register-row">
          <Form.Control
            type="datetime-local"
            value={nowStr}
            onChange={evt => {
              setNowStr(evt.target.value);
              console.log(evt.target.value);
            }}
          />
        </Form.Group>
        <Form.Group className="register-row">
          <Form.Control
            type="text"
            value={name}
            placeholder="지출 내역"
            onChange={evt => setName(evt.target.value)}
          />
        </Form.Group>
        <Form.Group className="register-row">
          <ButtonGroup>
            <TypeButtons
              active={activeBtn}
              setActiveBtn={setActiveBtn}
            />
          </ButtonGroup>
          <Col>
              <Form.Control
                type="number"
                value={price}
                placeholder="0"
                onChange={evt => setPrice(evt.target.value)}
              />
            </Col>
          <div>원</div>
        </Form.Group>
        <Form.Group className="register-row">
          결제 수단/할부
          <DropdownPaymentMethods
            gid={location.state.gid}
            active={activeMethod}
            setActiveMethod={setActiveMethod}
          />
          <Col>
            <Form.Control
              type="number"
              value={monthlyInstallment}
              placeholder="0"
              onChange={evt => setMonthlyInstallment(evt.target.value)}
            />
          </Col>
          <div>개월</div>
        </Form.Group>
        <Form.Group className="register-row">
          <Col>
            <Form.Control
              type="text"
              value={category}
              placeholder="카테고리"
              onChange={evt => setCategory(evt.target.value)}
            />
          </Col>
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

function _zeroPad(va) {
  return va.toString().padStart(2, '0');
}
