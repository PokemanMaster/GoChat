import React from 'react';
import ReactDOM from 'react-dom/client';
import Route from './route';
import './index.css';
import { BrowserRouter } from 'react-router-dom';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
        <BrowserRouter>
            <Route />
        </BrowserRouter>
    </React.StrictMode>
);
