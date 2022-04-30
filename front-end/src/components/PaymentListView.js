import './PaymentListView.css';
import SummaryBySign from './SummaryBySign';

function PaymentListView(props) {
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

  return (
    <article className="abs-payment-list-view">
      {paymentsByDate.length === 0 && <p>내역이 없어요.</p>}
      {paymentsByDate.map(payments =>
        <section
          key={payments.key}
          className="abs-payment-list-daily"
        >
          <header>
            <h5>{payments.key}</h5>
            <SummaryBySign
              payments={payments.items}
            />
          </header>
          <ul>
            {payments.items.map((item, i) =>
              <li
                className="abs-payment-item"
                key={`${item.date}: ${i}`}
              >
                {item.name}
                <span>{item.price.toLocaleString()}원</span>
              </li>)}
          </ul>
        </section>)}
    </article>
  );
}

export default PaymentListView;
