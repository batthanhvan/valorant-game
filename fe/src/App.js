import logo from './logo.svg';
import './App.css';
import { useEffect } from 'react';
import { useState } from 'react';
import SearchBar from './component/searchBar';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Routes,
  Link
} from "react-router-dom";
import Profile from './page/profile';
import Home from './page/Home';
import Report from './page/report';
import EditInfo from './page/EditInfo';
import Rep from './page/rep';
function App() {
  


  return (
    <Router>
   
    
    
      <Routes>
      {/* <nav className="App">
     
     <SearchBar />
     </nav> */}
        <Route path='/Profile' element={<Profile />}></Route>
        <Route path='/' element={<Home />}></Route>
        <Route path='/Report' element={<Report />}></Route>
        <Route path='/Edit' element={<EditInfo />}></Route>
        <Route path='/Rep' element={<Rep />}></Route>
      </Routes>
    
    
    </Router>
  );
}

export default App;
