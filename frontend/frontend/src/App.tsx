import React, { useEffect, useState } from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Borrow from "./components/Borrow";
import Borrowed from "./components/Borrowed";

// import UserCreate from "./components/UserCreate";
// import Login from "./components/Login";


 export default function App() {
  return (
    <Router>

   <div>

   <Navbar />

   <Routes>

       <Route path="/" element={<Borrow />} />

       <Route path="/Borrow" element={<Borrowed />} />
       
   </Routes>

   </div>

  </Router>

  );
 }
