import './index.css';
import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Home, Explore, ComingSoon, NotFound } from './routes';

const AppPage = () => {
  if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development')
    return <App />;
  else return <ComingSoon />;
};

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<AppPage />}>
          <Route index element={<Home />} />
          <Route path="explore" element={<Explore />} />
        </Route>
        <Route path="comingsoon" element={<ComingSoon />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
    <App />
  </React.StrictMode>
);
