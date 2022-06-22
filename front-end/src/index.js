import React from 'react';
import '@/index.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import App from '@/App';
import reportWebVitals from '@/reportWebVitals';
import { createRoot } from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

const container = document.getElementById('root');
const root = createRoot(container);
root.render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<App tab="home"/>} />
    </Routes>
  </BrowserRouter>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
