import React, { useEffect, useState } from "react";
import Axios from "axios";
import {  useNavigate, useLocation } from "react-router";
const Rep =() => {
  const url="http://localhost:8080/report"
    const location = useLocation();
    let navigate=useNavigate();
  const [data, setData] = useState({
      playerName: "",
      reportCategory: "",
      matchid: "",
      detail:""
  });
  function handleedit(e){
    
    Axios.get(url,{
      playerName: data.playerName,
      reportCategory: data.reportCategory,
      matchid: data.matchid,
      detail: data.detail
    })
    .then(res=>{
      console.log(res.data);
        
        if(res.data.username!==""){
            alert("Report Successful")
            navigate("/",{state:{user:location.state.token}});
        }
        else  alert("Report Failed")
       
    })
}
return(
    <div className="containerx">
  <form>
    
    <div className="group">      
      <input type="text" onChange={({target}) => setData(state => ({...state, playerName:target.value}))} value={data.playerName} required />
      <span className="highlight"  ></span>
      <span className="bar"></span>
      <label>playername</label>
    </div>
      
    <div className="group">      
      <input type="text" onChange={({target}) => setData(state => ({...state, tagline:target.value}))} value={data.tagline} required />
      <span className="highlight"></span>
      <span className="bar"></span>
      <label>reportCategory</label>
      
    </div>
    <div className="group">      
      <input type="text" onChange={({target}) => setData(state => ({...state, matchid:target.matchid}))} value={data.matchid} required />
      <span className="highlight"></span>
      <span className="bar"></span>
      <label>matchid</label>
     
    </div>
    <div className="group">      
      <input type="text" onChange={({target}) => setData(state => ({...state, detail:target.detail}))} value={data.detail} required />
      <span className="highlight"></span>
      <span className="bar"></span>
      <label>detail</label>
      <button  type="button" className='bt' onClick={() => handleedit()}>SEND</button>
    </div>
    
  </form>       
</div>
)
}
export default Rep;