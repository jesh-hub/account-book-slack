import '@/pages/PaymentRegisterView.scss';
import { Button, ButtonGroup, Col, Dropdown, DropdownButton, Form } from 'react-bootstrap';
import { useState } from 'react';
import useRequest from '@/common/useRequest';
import { useLocation } from 'react-router-dom';

const BtnTypes = [
  { key: 'income', uiText: '지출' },
  { key: 'outgoing', uiText: '수입' },
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

export default function PaymentRegisterView() {
  const location = useLocation();
  const [activeBtn, setActiveBtn] = useState(BtnTypes[0]);
  const [activeMethod, setActiveMethod] = useState();

  return (
    <article className="abs-register">
      <Form>
        <Form.Group className="register-row">
          <Form.Control
            type="date"
            placeholder="0"
          />
        </Form.Group>
        <Form.Group className="register-row">
          <Form.Control
            type="text"
            placeholder="지출 내역"
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
                placeholder="0"
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
              placeholder="0"
            />
          </Col>
          <div>개월</div>
        </Form.Group>
        <Form.Group className="register-row">
          <Col>
            <Form.Control
              type="text"
              placeholder="카테고리"
            />
          </Col>
        </Form.Group>
        <Button className="w-100">
          등록
        </Button>
      </Form>
    </article>
  );
}
