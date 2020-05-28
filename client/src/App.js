import React, { useState, useEffect } from 'react';
import server from './services/server'

function App() {

  useEffect(() => {
    async function getList() {
      let res = await server.ListQuestions();
    }
    getList()
  }, [])

  return (
    <div className="App">
      HELLO WORLD 1
    </div>
  );
}

export default App;
