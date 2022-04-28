import './App.css';
import {Dropdown, DropdownButton} from 'react-bootstrap';
import {useEffect, useReducer, useState} from 'react';
import PaymentListView from './components/PaymentListView';
import SummaryBySign from './components/SummaryBySign';
import * as Api from './common/Api';

function App() {
  const [currentMonth, setCurrentMonth] = useState(new Date().getMonth());

  const monthDropdownItems = [];
  for (let i = 0; i < 12; i++)
    monthDropdownItems.push(
      <Dropdown.Item
        key={i}
        disabled={currentMonth === i}
        onClick={() => setCurrentMonth(i)}
      >
        {i + 1}월
      </Dropdown.Item>);

  const [state, dispatch] = useReducer((_, action) => ({
    processing: action.processing,
    payments: action.data ?? []
  }), { processing: false, payments: [] });

  useEffect(() => {
    fetchData().then();

    async function fetchData() {
      const monthDateStr = ``+
        `${new Date().getFullYear()}-${String(currentMonth + 1).padStart(2, '0')}`;

      await Api.getWithDispatch(dispatch, '/payments', {
        start: monthDateStr,
        end: monthDateStr
      });
    }
  }, [currentMonth]);

  return (
    <div className="app">
      <header className="app-header">
        <DropdownButton
          className="abs-full-width-button"
          variant="outline-primary"
          title={`${currentMonth + 1}월`}
        >
          {monthDropdownItems}
        </DropdownButton>
        <article className="abs-monthly-summary">
          <SummaryBySign
            payments={state.payments}
            processing={state.processing}
            mt="1"
          />
        </article>
      </header>
      <main className="app-main">
        <PaymentListView payments={state.payments} />
      </main>
    </div>
  );
}

export default App;
