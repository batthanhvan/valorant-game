import SearchBar from "../component/searchBar";
import './Home.css';
import Axios from "axios";
import React, { useEffect, useState } from "react";
import {  useNavigate, useLocation } from "react-router";
import Leaderboard from "./leaderboard";
import Report from "./report";
const Home =() => {
  let navigate=useNavigate();
  const location = useLocation();
  const url="http://localhost:8080/login";
  const url1="http://localhost:8080/register";
  const [data, setData] = useState({
    username: "",
    password: ""
});
const [user, setUser] = useState("")
useEffect(() => {
  if(location.state){
    setUser(location.state.user)
  }
} ,[])
function submit(e){
    e.preventDefault();
    Axios.post(url,{
      username: data.username,
      password: data.password  
    })
    .then(res=>{
      console.log(res.data);
        setUser(res.data)
        if(res.data.username!==""){
            alert("Login Successful")
            console.log(user)
        }
        else  alert("Login Failed")
       
    })
}
function submit1(e){
  e.preventDefault();
  Axios.post(url1,{
    username: data.username,
    password: data.password  
  })
  .then(res=>{
    
    
      if(res.data.username=="registration success"){
          alert("Register Failed")
      }
      else alert("Register Successful ")
      console.log(user)
  })
}
function handle(e){
const newdata={...data};
newdata[e.target.value]=e.target.value;
setData(newdata);
console.log(newdata)
}
function handleEdit() {
        
  
  navigate("/Edit",{state:{token:user}});
}
function handleLogout() {
  setUser("")
}
function ShowReport() {
  navigate("/Report",{state:{token:user}});
}
function Report() {
  navigate("/Rep",{state:{token:user,username:data.username}});
}
if(user!==""){
    return( <div className="Container1">
        <header>
        <button  type="button" className='bt' onClick={() => handleEdit()}>EDIT</button>
        <button  type="button" className='bt' onClick={() => handleLogout()}>LOGOUT</button>
        <button  type="button" className='bt' onClick={() => ShowReport()}>HISTORY</button>
        <button  type="button" className='bt' onClick={() => Report()}>REPORT</button>
        </header>
        
        <div className="da">
          <Leaderboard></Leaderboard>
          </div>
        </div>
    );}
    return (<div  className="con">
      <div className="header1">
  <div className="login-container">
    <form >
      <input type="text" placeholder="Username" name="username" onChange={({target}) => setData(state => ({...state,username:target.value}))} value={data.username} />
      <input type="text" placeholder="Password" name="password"onChange={({target}) => setData(state => ({...state,password:target.value}))} value={data.password} />
      <button type="submit" onClick={(e)=>submit(e)}>Login</button>
      <button type="submit" onClick={(e)=>submit1(e)}>Register</button>
    </form>
    </div>
    
    </div>

      <div className="Container2">
      <SearchBar className="ip" />
      </div>
      </div>
    );
}
export default Home;