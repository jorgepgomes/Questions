import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import "./global.css";
import "./App.css";

import Routes from './routes'

export default function App() {
  return (
    <div id="app">
    <Router>
          <Routes/>
    </Router>
    </div>
  );
}
