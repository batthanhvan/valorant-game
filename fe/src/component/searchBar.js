import React, { useEffect, useState } from "react";
import './searchBar.css';
import { useNavigate } from 'react-router';
const SearchBar = () => {
    let navigate=useNavigate();
    const [state, setState] = React.useState({
        firstName: ""
      })
      
      function handleChange(evt) {
          
        const value = evt.target.value;
        setState({
          firstName:evt.target.value
          
        });
        console.log(state.firstName);
      }
      function handleSubmit() {
        
        console.log(state);
        navigate("/Profile",{state:{firstName:state.firstName}});
      }
      function handleReport() {
        
        console.log(state);
        navigate("/Report",{state:{firstName:state.firstName}});
      }
      return (
    <div className="searchBar">
    <form action="/" method="get">
        <label htmlFor="header-search">
            <span className="visually-hidden">Search blog posts</span>
        </label>
        <div className='modal'>
        
        <input className='ip'
            type="text"
            id="header-search"
            placeholder=""
            name="s" 
            onChange={handleChange}
        />
        <button  type="button" className='bt' onClick={() => handleSubmit()}>Search</button>
        
        </div>
        
    </form>
    </div>
      );
      }

export default SearchBar;