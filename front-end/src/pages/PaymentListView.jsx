import '@/pages/PaymentListView.scss';
import { BsExclamationCircle } from 'react-icons/bs';
import { RiEdit2Line } from 'react-icons/ri';
import { Button, Dropdown } from 'react-bootstrap';
import { useState } from 'react';
import { useLocation } from 'react-router-dom';
import SummaryBySign from '@/components/SummaryBySign';
import useRequest from '@/common/useRequest';
import ProcessingSpinner from '@/common/ProcessingSpinner';

const CurrentYear = 2022;

function _buildDateRange(year, month) {
  return {
    dateFrom: `${year}-${String(month + 1).padStart(2, '0')}`,
    dateTo: `${year}-${String(month + 2).padStart(2, '0')}`
  }
}

function korStdTimeStrToDateStr(date) {
  const _date = new Date(date);
  return [
    _date.getFullYear(),
    `${_date.getMonth() + 1}`.padStart(2, '0'),
    `${_date.getDate()}`.padStart(2, '0')
  ].join('-');
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
              <span className="li__title">{item.name}</span>
              <span className="li__price">{item.price.toLocaleString()}원</span>
              <Button
                size="xs"
                variant="soft-clear"
                className="li__edit-btn"
              ><RiEdit2Line /></Button>
            </li>)
        }
      </ul>
    </section>);
}

export default function PaymentListView() {
  const location = useLocation();
  const [currentMonth, setCurrentMonth] = useState(new Date().getMonth());

  const [processing, payments] = useRequest(
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
            disabled={processing}
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
        {
          processing ?
            <ProcessingSpinner processing={processing} /> :
            <PaymentList payments={payments} />
        }
      </section>
    </article>
  );
}
