import React, { useEffect, useState } from "react";
import {  useNavigate, useLocation } from "react-router";
import './report.css';
import axios from "axios";
const Report =() => {
    const location = useLocation();
    const [report, setReport] = useState([]);
    const navigate=useNavigate();
    useEffect   (() => {
      const fetchReport = async () => {
    
        let token=(location.state.token.token);
        console.log(token);
        const url="http://localhost:8080/admin/reports/"+"an";
          axios.get(url,null,{
           params: {
              token
            }
           })
          .then(res=>{
            console.log(res.data.success)  
            if(res.data.success!="false"){
             
            }
           else alert("Edit Failed")
           
      })
    }
      fetchReport(); 
      
  },[]);
  for(let rep in report){
    console.log(report[rep]);
  }
  function handleBack(){
    navigate("/",{state:{user:location.state.token}});
  }
    return(
      <div>
<table striped bordered hover variant="dark">
  <thead>
    <tr>
      <th>reportCategory</th>
      <th>reportDetail</th>
      <th>reportDate</th>
      <th>recordLink</th>
    </tr>
  </thead>
  <tbody>
     
    <tr>
      <td>Negative Attitude</td>
      <td>YOXEBvmMjRtyv2G8hBLUMD7bkjSvFRJrvA3uSFZm296Kl2XcLL8Mu7YwmraZRqAbkj9YX</td>
      <td>2022-01-07</td>
      <td>NXWOuLSptdOfpKNiBX9zngNXXPc9iQ</td>
    </tr>
    <tr>
      <td>Assisting Enemy</td>
      <td>y2iAVZmTgTjXXJjh3ErDXsad8qNbFlWRQ1JmHW7NkgfFmMepwNkblHoke3sIZdxsVPP1Rof3vj4TpBeBo4yI67RS</td>
      <td>2021-08-27</td>
      <td>5WWDgHB9J8UJBGzpZW50L4HN5QggZbhKsx78kqN</td>
    </tr>
    <tr>
      <td>Assisting Enemy</td>
      <td>tdxkXv5VoFffzX1ncs3i3EAKv</td>
      <td>2021-02-17</td>
      <td>Yz9gRGCO5yUzFxHfpknkWz85Ll9ZB0bNCh8zTRwkyfvNdUEGZSF9D36idBQyM</td>
    </tr>
    <tr>
      <td>Negative Attitude</td>
      <td>X95qeiz5UMuA0ZItBnDf6YnSTLxBqyJGq9YPzjhQVXVbqTzkmOTBkt</td>
      <td>2020-02-28</td>
      <td>zhlIj4kqhsRGQMb</td>
    </tr>
  </tbody>
</table>
<div className="ass">
<button onClick={() => handleBack()}>back Home</button>
</div>
</div>
    );
}
export default Report;