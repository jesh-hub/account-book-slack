import '@/pages/PaymentListView.scss';
import { BsExclamationCircle } from 'react-icons/bs';
import { Dropdown } from 'react-bootstrap';
import { useState } from 'react';
import { useLocation } from 'react-router-dom';
import SummaryBySign from '@/components/SummaryBySign';
import useRequest from '@/common/useRequest';

const CurrentYear = 2022;

function _buildDateRange(year, month) {
  return {
    dateFrom: `${year}-${String(month + 1).padStart(2, '0')}`,
    dateTo: `${year}-${String(month + 2).padStart(2, '0')}`
  }
}

function korStdTimeStrToDateStr(date) {
  const _date = new Date(date);
  return `${_date.getFullYear()}-${`${_date.getMonth() + 1}`.padStart(2, '0')}-${`${_date.getDate()}`.padStart(2, '0')}`;
}

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
  let curDateStr;
  const paymentsByDate = props.payments.sort((a, b) => a.date > b.date ? -1 : 1) // 내림차순
    .reduce((acc, /** @type {Payment} */ cur) => {
      const dateStr = korStdTimeStrToDateStr(cur.date);
      if (dateStr !== curDateStr) {
        acc.push({
          key: cur.id,
          date: dateStr,
          items: []
        });
        curDateStr = dateStr;
      }
      acc[acc.length - 1].items.push(cur);
      return acc;
    }, []);
  if (paymentsByDate.length === 0)
    return <p><BsExclamationCircle />내역이 없어요.</p>;
  return paymentsByDate.map(payments =>
    <section
      key={payments.key}
      className="daily-payments"
    >
      <header>
        <h5>{payments.date}</h5>
        <SummaryBySign payments={payments.items} />
      </header>
      <ul>
        {
          payments.items.map((item, i) =>
            <li key={`${payments.key}: ${i}`}>
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

  const [_processing, payments] = useRequest(
    '/v1/payment', {
      groupId: location.state.gid,
      ..._buildDateRange(CurrentYear, currentMonth)
    }, [currentMonth], location.state.payments || []);

  return (
    <article className="abs-payments">
      <header className="header">
        <Dropdown>
          <Dropdown.Toggle
            variant="outline-primary"
            className="w-100"
            disabled
          >{CurrentYear}년</Dropdown.Toggle>
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
        <section className="summary">
          <SummaryBySign
            payments={payments}
            className="bilateral-align"
          />
        </section>
      </header>
      <section className="payments">
        <PaymentList payments={payments} />
      </section>
    </article>
  );
}
