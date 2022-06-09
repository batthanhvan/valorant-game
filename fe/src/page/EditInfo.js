import React, { useEffect, useState } from "react";
import Axios from "axios";
import {  useNavigate, useLocation } from "react-router";
import './Editname.css'
const EditInfo =() => {
  const location = useLocation();
  let navigate=useNavigate();
const [data, setData] = useState({
    playerName: "",
    tagline: ""
});


function handleedit(e){
 let token=(location.state.token.token);
  console.log(token);
  const url="http://localhost:8080/players/modify/"+data.playerName+"/"+data.tagline;
    Axios.post(url,null,{
     params: {
        token
      }
     })
    .then(res=>{
      console.log(res.data.success)
      if(res.data.success!="false"){
        alert("Edit Successful")
        navigate("/",{state:{user:location.state.token}});
      }
     else alert("Edit Failed")

    })
  //   if(1===1){
  //     alert("Edit Successful")
  //     navigate("/");
  // }
  
  // else  alert("Edit Failed")

}


    return(
      <div className="big-container">
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
      <label>tagline</label>
      <button  type="button" className='bt' onClick={() => handleedit()}>EDIT</button>
    </div>
    
  </form>       
</div>
</div>
            );
}
export default EditInfo

