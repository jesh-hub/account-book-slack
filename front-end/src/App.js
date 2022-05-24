import '@/App.css';
import {Dropdown, DropdownButton} from 'react-bootstrap';
import {useState} from 'react';
import PaymentListView from '@/components/PaymentListView';
import ProcessingSpinner from '@/common/ProcessingSpinner';
import SummaryBySign from '@/components/SummaryBySign';
import useRequest from '@/common/useRequest';
import ErrorLogger from '@/components/ErrorLogger';
import ErrorContext from '@/common/ErrorContext';

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

  const start = `` +
    `${new Date().getFullYear()}-${String(currentMonth + 1).padStart(2, '0')}`;
  const end = `` +
    `${new Date().getFullYear()}-${String(currentMonth + 1).padStart(2, '0')}`;
  const [processing, payments] = useRequest(
    '/payments', { start, end }, [currentMonth], []);

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
            payments={payments}
            mt="1"
          />
          <ProcessingSpinner processing={processing} />
        </article>
      </header>
      <main className="app-main">
        {!processing && <PaymentListView payments={payments} />}
      </main>
      <aside className="app-aside">
        <ErrorContext.Consumer>
          {
            value =>
              <ErrorLogger
                errors={value.errors}
                deleteError={value.deleteError}
              />
          }
        </ErrorContext.Consumer>
      </aside>
    </div>
  );
}

export default App;
