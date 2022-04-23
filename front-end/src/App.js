import './App.css';
import {Dropdown, DropdownButton} from 'react-bootstrap';
import {useState} from 'react';
import SummaryBySign from './components/SummaryBySign';

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
        <SummaryBySign
          currentMonth={currentMonth}
          mt="1"
        />
      </header>
    </div>
  );
}

export default App;
