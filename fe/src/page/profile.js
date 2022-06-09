import React, { useEffect, useState } from "react";
import { useLocation } from "react-router";
import './profile.css';
import img from './cropped-150-150-1187222.png'
const Profile = () => {
  const location = useLocation();
    const [match, setMatch] = useState([]);
    const[playerInfo, setPlayerInfo] = useState([]);
    useEffect   (() => {
      const fetchMatch = async () => {
        let response = await fetch("http://localhost:8080/players/"+location.state.firstName);
        response = await response.json()
        setPlayerInfo(response.data);
        let res=await fetch("http://localhost:8080/matches/"+location.state.firstName);
        res=await res.json(); 
        setMatch(res.data.result);
       
      }
      fetchMatch(); 
      
  },[]);
  for ( let mat in match){
    console.log(match[mat]);
  }
  const matches=match.slice(0,9);
  
  const pInfo=Object.entries(playerInfo);
console.log(pInfo);
  // for(let info in playerInfo){
  //   console.log(playerInfo[info]);}
  return (
    <div className="container">
      <header>
        
      </header>
  <aside>
   <div>
     {pInfo.map(info => (
    <div>
   <div className="basic-info">
    <div className="basic-info-container">
          <img className="rankImg" src="https://d2j6dbq0eux0bg.cloudfront.net/images/224443/1603512155.jpg"></img>
          <h1>{info[1].playerName}</h1>
          </div>
          <div className="player-info-container">
            <h1 className="tag">Wins</h1>
            <h1 className="tag2">{info[1].wins}</h1>
             </div>
             <hr  width="75%" size="5px" align="left" color="red" className="line"  ></hr>
        </div>
        <div className="player-info-container">
            <h1 className="tag">Kills</h1>
            <h1 className="tag2">{info[1].kills}</h1>
             </div>
             <hr  width="75%" size="5px" align="left" color="red" className="line"  ></hr>
             <div className="player-info-container">
            <h1 className="tag">Asist</h1>
            <h1 className="tag2">{info[1].assists}</h1>
             </div>
             <hr  width="75%" size="5px" align="left" color="red" className="line"  ></hr>
             <div className="player-info-container">
            <h1 className="tag">Kda</h1>
            <h1 className="tag2">{info[1].killsPerRound.toFixed(2)}</h1>
             </div>
             <hr  width="75%" size="5px" align="left" color="red" className="line"  ></hr>
             <div className="player-info-container">
            <h1 className="tag">First </h1>
            <h1 className="tag2">{info[1].firstBloods}</h1>
             </div>
             <hr  width="75%" size="5px" align="left" color="red" className="line"  ></hr>
             <div className="player-info-container">
            <h1 className="tag">Aces</h1>
            <h1 className="tag2">{info[1].aces}</h1>
             </div>
             <hr  width="75%" size="5px" align="left" color="red" className="line"  ></hr>
             <div className="player-info-container">
            <h1 className="tag">Clutch</h1>
            <h1 className="tag2">{info[1].clutches}</h1>
             </div>
             <hr  width="75%" size="5px" align="left" color="red" className="line"  ></hr>
             <div className="player-info-container">
            <h1 className="tag">Mvp</h1>
            <h1 className="tag2">120</h1>
             </div>
             <hr  width="75%" size="5px" align="left" color="red" className="line"  ></hr>
             <div className="player-info-container">
            <h1 className="tag">HS</h1>
            <h1 className="tag2">120</h1>
             </div>
             <hr  width="75%" size="5px" align="left" m color="red" className="line"  ></hr>
             </div>
             ))
             }
             </div>
  </aside>
  <main>
   <div>
    {matches.map(matf => (
<div className="card">
  <div className="card-container">
  <div className="match__portrait">
  <img className="avatar" src={require('./cropped-150-150-1187222-removebg-preview.png')} />
    </div>
    <div className="center">
    <div className="match__detail">
      <h1 className="tex">{matf.matchServer}</h1>
      </div>
      </div>
      <div className="center">
    <div className="match__result">
    <h1 className="tex">{matf.mapName}</h1>
    </div>
    </div>
    <div className="center"></div>
    <div className="match__row-stats">
    <h1 className="tex">{matf.modeName}</h1>
    </div>
    
</div>
</div>))}
</div>
  </main>
  
     </div>
  );
  

}
export default Profile