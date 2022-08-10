import '@/pages/PaymentRegisterView.scss';
import { Button, ButtonGroup, Col, Dropdown, DropdownButton, Form } from 'react-bootstrap';
import { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import axios from 'axios';
import useRequest from '@/common/useRequest';

const BtnTypes = [
  { key: 'income', uiText: '지출', value: -1 },
  { key: 'outgoing', uiText: '수입', value: 1 },
];
const DateParamSuffix = ':00+09:00';

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
  const [isFirst, setIsFirst] = useState(true);
  const [isWaitingPaymentMethods, paymentMethods] = useRequest(
    '/v1/paymentMethod', { groupId: props.gid }, [], []);

  useEffect(() => {
    if (isFirst && ! isWaitingPaymentMethods)
      setIsFirst(false);
    else if (! isWaitingPaymentMethods)
      props.onInit(paymentMethods);
  }, [isWaitingPaymentMethods]);

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
  const navigate = useNavigate();

  const { prev } = location.state;
  const [date, setDate] = useState(prev !== undefined ? new Date(prev.date) : new Date());
  const [dateStr, setDateStr] = useState(dateToDateStr(date));
  const [name, setName] = useState(prev?.name || '');
  const [activeBtn, setActiveBtn] = useState(getInitialPaymentType());
  const [price, setPrice] = useState(Math.abs(prev?.price) || '');
  const [activeMethod, setActiveMethod] = useState();
  const [monthlyInstallment, setMonthlyInstallment] = useState(prev?.monthlyInstallment || '');
  const [category, setCategory] = useState(prev?.category || '');

  function getInitialPaymentType() {
    return prev === undefined ? BtnTypes[0] :
      BtnTypes.find(type => prev.price / Math.abs(prev.price) === type.value);
  }

  async function submit(arg) {
    arg.preventDefault();
    try {
      const register = prev !== undefined ? axios.put : axios.post;
      await register(`${process.env.REACT_APP_ABS}/v1/payment` +
        (prev !== undefined ? `/${prev.id}` : ''), {
        date: dateStr + DateParamSuffix,
        // date: date.getTime(),
        name,
        category,
        price: price * activeBtn.value,
        monthlyInstallment: +monthlyInstallment || 0,
        paymentMethodId: activeMethod?.id,
        groupId: location.state.gid,
        regUserId: props.userInfo.id
      });
      navigate(-1);
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
            value={dateStr}
            onChange={evt => {
              setDateStr(evt.target.value);
              setDate(dateStrToDate(evt.target.value));
            }}
            required
          />
        </Form.Group>
        <Form.Group className="register-row">
          <Form.Control
            type="text"
            value={name}
            placeholder="지출 내역"
            onChange={evt => setName(evt.target.value)}
            required
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
            onInit={(paymentMethods) => {
              if (prev !== undefined) {
                const method = paymentMethods.find(method => method.id === prev.paymentMethodId);
                setActiveMethod(method);
              }
            }}
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
              required
            />
          </Col>
        </Form.Group>
        <Button
          className="w-100"
          type="submit"
          disabled={! activeBtn || ! activeMethod}
        >
          등록
        </Button>
      </Form>
    </article>
  );
}

function dateToDateStr(date) {
  return `` +
    `${date.getFullYear()}-${_zeroPad(date.getMonth() + 1)}-${_zeroPad(date.getDate())}` +
    `T${_zeroPad(date.getHours())}:${_zeroPad(date.getMinutes())}`;
}

function dateStrToDate(dateStr) {
  const str = dateStr.replace(/\D/g, '');
  return new Date(
    +str.substring(0, 4),
    +str.substring(4, 6) - 1,
    +str.substring(6, 8),
    +str.substring(8, 10),
    +str.substring(10, 12),
    +str.substring(12, 14));
}

function _zeroPad(va) {
  return va.toString().padStart(2, '0');
}
