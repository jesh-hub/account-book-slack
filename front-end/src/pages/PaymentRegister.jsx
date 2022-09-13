import '@/pages/PaymentRegister.scss';
import { Button, ButtonGroup, Col, Dropdown, DropdownButton, Form } from 'react-bootstrap';
import { useCallback, useEffect, useMemo, useReducer, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import axios from 'axios';
import useRequest from '@/common/useRequest';
import ProcessingSubmitButton from '@/common/ProcessingSubmitButton';
import { getDateStr } from '@/common/DateUtil';
import { doPostRequest, doPutRequest } from '@/common/Api';

const PaymentTypes = [
  { key: 'income', uiText: '지출', value: -1 },
  { key: 'outgoing', uiText: '수입', value: 1 },
];
const DateParamSuffix = ':00+09:00';

function PaymentTypeRadio({ type, setType }) {
  return PaymentTypes.map(paymentType =>
    <Button
      variant="outline-primary"
      size="sm"
      key={paymentType.key}
      active={type === paymentType}
      onClick={() => setType(paymentType)}
    >
      {paymentType.uiText}
    </Button>);
}

function DropdownPaymentMethods(props) {
  const [isFirst, setIsFirst] = useState(true);
  const [isWaitingPaymentMethods, paymentMethods] = useRequest(
    `/v1/group/${props.gid}/paymentMethod`, null, [], []);

  useEffect(() => {
    if (isFirst && ! isWaitingPaymentMethods)
      setIsFirst(false);
    else if (! isWaitingPaymentMethods)
      props.onInit(paymentMethods);
  // eslint-disable-next-line
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

const ACTION_TYPE = {
  PAYMENT_TYPE: 0,
};

function reducer(formData, action) {
  const { type, name, value } = action;
  switch (type) {
    case ACTION_TYPE.PAYMENT_TYPE:
    {
      const ret = {
        ...formData,
        paymentType: action.value,
      };
      if (action.value.value > 0)
        ret.monthlyInstallment = '';
      return ret;
    }
    default:
      return {
        ...formData,
        [name]: value,
      };
  }
}

export default function PaymentRegister({ userInfo }) {
  const navigate = useNavigate();

  const location = useLocation();
  const { prev, gid } = location.state;

  const [processing, setProcessing] = useState(false);
  const initialFormData = useMemo(() => ({
    date: getDateStr(prev !== undefined ? new Date(prev.date) : new Date(), 'T'),
    name: prev?.name || '',
    monthlyInstallment: prev?.monthlyInstallment || '',
    price: Math.abs(prev?.price) || '',
    paymentType: prev === undefined ? PaymentTypes[0] :
      PaymentTypes.find(type => prev.price / Math.abs(prev.price) === type.value),
    category: prev?.category || '',
    paymentMethod: undefined,
  }), [prev]);
  const [formData, dispatch] = useReducer(reducer, initialFormData, undefined);

  async function submit(arg) {
    arg.preventDefault();
    try {
      setProcessing(true);
      const params = {
        date: formData.date + DateParamSuffix,
        // date: date.getTime(),
        name: formData.name,
        category: formData.category,
        price: formData.price * formData.paymentType.value,
        monthlyInstallment: +formData.monthlyInstallment,
        paymentMethodId: formData.paymentMethod?.id,
        groupId: gid,
      };

      if (prev === undefined) {
        await doPostRequest(`/v1/group/${gid}/payment`, {
          ...params,
          regUserId: userInfo.id,
        });
      } else {
        await doPutRequest(`/v1/payment/${prev.id}`, {
          ...params,
          modUserId: userInfo.id
        });
      }
      setProcessing(false);
      navigate(-1);
    } catch (e) {
      console.log(e);
      setProcessing(false);
    }
  }

  const handleFormDataDefaultChanged = useCallback(evt => {
    const { name, value } = evt.target;
    dispatch({
      name,
      value,
    });
  }, []);

  const handleFormDataTypeChanged = useCallback(type => dispatch({
    type: ACTION_TYPE.PAYMENT_TYPE,
    value: type,
  }), []);

  const handleFormDataMethodChanged = useCallback(value => dispatch({
    name: 'paymentMethod',
    value,
  }), []);

  const handleMethodInitialized = useCallback(methods => {
    if (prev === undefined)
      return;
    handleFormDataMethodChanged(
      methods.find(method => method.id === prev.paymentMethodId));
  }, [prev, handleFormDataMethodChanged]);


  return (
    <article className="abs-register">
      <Form onSubmit={submit}>
        <Form.Group className="register-row">
          <Form.Control
            name="date"
            type="datetime-local"
            value={formData.date}
            onChange={handleFormDataDefaultChanged}
            required
          />
        </Form.Group>
        <Form.Group className="register-row">
          <Form.Control
            name="name"
            type="text"
            value={formData.name}
            placeholder="지출 내역"
            onChange={handleFormDataDefaultChanged}
            required
          />
        </Form.Group>
        <Form.Group className="register-row">
          <ButtonGroup>
            <PaymentTypeRadio
              type={formData.paymentType}
              setType={handleFormDataTypeChanged}
            />
          </ButtonGroup>
          <Col>
              <Form.Control
                name="price"
                type="number"
                value={formData.price}
                placeholder="0"
                onChange={handleFormDataDefaultChanged}
              />
            </Col>
          <div>원</div>
        </Form.Group>
        <Form.Group className="register-row">
          결제 수단/할부
          <DropdownPaymentMethods
            gid={location.state.gid}
            active={formData.paymentMethod}
            setActiveMethod={handleFormDataMethodChanged}
            onInit={handleMethodInitialized}
          />
          <Col>
            <Form.Control
              name="monthlyInstallment"
              type="number"
              value={formData.monthlyInstallment}
              placeholder="0"
              disabled={formData.paymentType?.value > 0}
              onChange={handleFormDataDefaultChanged}
            />
          </Col>
          <div>개월</div>
        </Form.Group>
        <Form.Group className="register-row">
          <Col>
            <Form.Control
              name="category"
              type="text"
              value={formData.category}
              placeholder="카테고리"
              onChange={handleFormDataDefaultChanged}
              required
            />
          </Col>
        </Form.Group>
        <ProcessingSubmitButton
          processing={processing}
          disabled={! formData.paymentType || ! formData.paymentMethod}
        >
          등록
        </ProcessingSubmitButton>
      </Form>
    </article>
  );
}
