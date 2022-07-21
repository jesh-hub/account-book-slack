import '@/pages/PaymentListView.scss';
import { Dropdown } from 'react-bootstrap';
import { useState } from 'react';

function MonthDropdownItems(props) {
  return new Array(12).fill(undefined).map((_, i) =>
    <Dropdown.Item
      key={i}
      disabled={props.currentMonth === i}
      href="#"
      onClick={() => props.setCurrentMonth(i)}
    >
      {i + 1}월
    </Dropdown.Item>)
}

export default function PaymentListView() {
  const [currentMonth, setCurrentMonth] = useState(new Date().getMonth());

  return (
    <article className="abs-payments">
      <header>
        <Dropdown>
          <Dropdown.Toggle variant="outline-primary" className="w-100">
            {currentMonth + 1}월
          </Dropdown.Toggle>
          <Dropdown.Menu className="w-100">
            <MonthDropdownItems
              currentMonth={currentMonth}
              setCurrentMonth={setCurrentMonth}
            />
          </Dropdown.Menu>
        </Dropdown>
      </header>
    </article>
  );
}
