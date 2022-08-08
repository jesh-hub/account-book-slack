import '@/pages/PaymentListView.scss';
import { AiOutlinePlus } from 'react-icons/ai';
import { BsExclamationCircle } from 'react-icons/bs';
import { Button, Dropdown } from 'react-bootstrap';
import { useState } from 'react';
import { useLocation } from 'react-router-dom';
import SummaryBySign from '@/components/SummaryBySign';

function MonthDropdownItems(props) {
  return new Array(12).fill(undefined).map((_, i) =>
    <Dropdown.Item
      key={i}
      disabled={props.currentMonth === i}
      href="#"
      onClick={() => props.setCurrentMonth(i)}
    >
      {i + 1}월
    </Dropdown.Item>);
}

function PaymentList(props) {
  let currentDate;
  const paymentsByDate = props.payments.sort((a, b) => a.date > b.date ? -1 : 1) // 내림차순
    .reduce((acc, /** @type {Payment} */ cur) => {
      if (cur.date !== currentDate) {
        acc.push({ key: cur.date, items: [] });
        currentDate = cur.date;
      }
      acc[acc.length - 1].items.push(cur);
      return acc;
    }, []);
  if (paymentsByDate.length === 0)
    return <p><BsExclamationCircle/>내역이 없어요.</p>;
  return paymentsByDate.map(payments =>
    <section
      key={payments.key}
      className="daily-payments"
    >
      <header>
        <h5>{payments.key}</h5>
        <SummaryBySign
          payments={payments.items}
        />
      </header>
      <ul>
        {
          payments.items.map((item, i) =>
            <li key={`${item.date}: ${i}`}>
              <span>{item.name}</span>
              <span>{item.price.toLocaleString()}원</span>
            </li>)
        }
      </ul>
    </section>);
}

export default function PaymentListView() {
  const location = useLocation();
  const [currentMonth, setCurrentMonth] = useState(new Date().getMonth());

  return (
    <article className="abs-payments">
      <header className="header">
        <Dropdown>
          <Dropdown.Toggle
            variant="outline-primary"
            className="w-100"
            disabled
          >2022년</Dropdown.Toggle>
        </Dropdown>
        <Dropdown>
          <Dropdown.Toggle
            variant="outline-primary"
            className="w-100"
          >{currentMonth + 1}월</Dropdown.Toggle>
          <Dropdown.Menu className="w-100">
            <MonthDropdownItems
              currentMonth={currentMonth}
              setCurrentMonth={setCurrentMonth}
            />
          </Dropdown.Menu>
        </Dropdown>
      </header>
      <section className="summary">
        <SummaryBySign
          payments={location.state.payments || []}
          className="bilateral-align"
        />
      </section>
      <section className="payments">
        <PaymentList payments={location.state.payments || []} />
      </section>
      <section className="register-btn">
        <Button
          className="w-100"
          variant="outline-primary"
          disabled
        ><AiOutlinePlus /></Button>
      </section>
    </article>
  );
}
