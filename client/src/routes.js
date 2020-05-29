import React from 'react';
import {Routes, Route } from 'react-router-dom';

import Home from './pages/home/index';
import Details from './pages/details/index';

export default function MainRoutes(){
    return (
        <Routes>
            <Route path="/" element={<Home/>}/>
            <Route path="/:id" element={<Details/>}/>
        </Routes>
    )
}
