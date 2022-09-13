import '@/pages/PaymentList.scss';
import { BsExclamationCircle } from 'react-icons/bs';
import { RiEdit2Line } from 'react-icons/ri';
import { Button, Dropdown } from 'react-bootstrap';
import { useCallback, useMemo, useState } from 'react';
import { useLocation } from 'react-router-dom';
import SummaryBySign from '@/components/SummaryBySign';
import ProcessingSpinner from '@/common/ProcessingSpinner';
import { useGetRequest } from '@/common/Api';
import { getDateDateStr, getDateMonthStr } from '@/common/DateUtil';
import useRouterNavigateWith from '@/common/useRouterNavigateWith';

function MonthDropdownItems({ curMonth, setCurMonth }) {
  const months = useMemo(() =>
    [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11], []);

  return months.map(month =>
    <Dropdown.Item
      key={month}
      disabled={curMonth === month}
      href="#"
      onClick={() => setCurMonth(month)}
    >
      {month + 1}월
    </Dropdown.Item>);
}

function DailyPaymentList({ payments }) {
  const navigateWith = useRouterNavigateWith();

  let curDateStr;
  const paymentsByDate = payments.reduce((acc, /** @type {Payment} */ cur) => {
    const dateStr = getDateDateStr(new Date(cur.date));
    if (dateStr !== curDateStr) {
      acc.push({
        id: cur.id,
        dateDateStr: dateStr,
        items: []
      });
      curDateStr = dateStr;
    }
    acc[acc.length - 1].items.push(cur);
    return acc;
  }, []);

  if (paymentsByDate.length === 0)
    return <p><BsExclamationCircle />내역이 없어요.</p>;
  return paymentsByDate.map(dailyPayments =>
    <section
      key={dailyPayments.id}
      className="daily-payments"
    >
      <header>
        <h5>{dailyPayments.dateDateStr}</h5>
        <SummaryBySign payments={dailyPayments.items} />
      </header>
      <ul>
        {
          dailyPayments.items.map(item =>
            <li key={item.id}>
              <span className="li__title">{item.name}</span>
              <span className="li__price">{item.price.toLocaleString()}원</span>
              <Button
                size="xs"
                variant="soft-clear"
                className="li__edit-btn"
                onClick={() => {
                  navigateWith('/payments/register', {
                    gid: item.groupId,
                    prev: item,
                  });
                }}
              ><RiEdit2Line /></Button>
            </li>)
        }
      </ul>
    </section>);
}

export default function PaymentList() {
  const location = useLocation();
  const { gid } = location.state;

  const [curDate, setCurDate] = useState(new Date());
  const [curYearDisplayed, curMonthDisplayed] = useMemo(() =>
    [`${curDate.getFullYear()}년`, `${curDate.getMonth() + 1}월`], [curDate]);
  const curMonthValue = useMemo(() => curDate.getMonth(), [curDate]);

  const dateParams = useMemo(() => {
    const dateFrom = new Date(curDate.getTime()),
      dateTo = new Date(curDate.getTime());
    dateTo.setMonth(dateTo.getMonth() + 1);
    return {
      dateFrom: getDateMonthStr(dateFrom),
      dateTo: getDateMonthStr(dateTo)
    };
  }, [curDate]);
  const [payments, processing] = useGetRequest(`/v1/group/${gid}/payment`, dateParams);

  const handleChangeCurMonth = useCallback((month) => {
    setCurDate(curDate => {
      const date = new Date(curDate.getTime());
      date.setMonth(month);
      return date;
    });
  }, []);

  return (
    <article className="abs-payment-list">
      <header className="header">
        <Dropdown>
          <Dropdown.Toggle
            variant="outline-primary"
            className="w-100"
            disabled
          >{curYearDisplayed}</Dropdown.Toggle>
        </Dropdown>
        <Dropdown>
          <Dropdown.Toggle
            variant="outline-primary"
            className="w-100"
            disabled={processing}
          >{curMonthDisplayed}</Dropdown.Toggle>
          <Dropdown.Menu className="w-100">
            <MonthDropdownItems
              curMonth={curMonthValue}
              setCurMonth={handleChangeCurMonth}
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
            <DailyPaymentList payments={payments} />
        }
      </section>
    </article>
  );
}
