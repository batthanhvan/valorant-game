import React, { useEffect, useState } from "react";
import Axios from "axios";
import { useLocation } from "react-router";
import './leaderboard.css';
const Leaderboard = () => {

    return(
        <div className="container1">
		<div className="leaderboard">
			<div className="head">
				<i className="fas fa-crown"></i>
				<h1>TOP PLAYERS</h1>
			</div>
			<div className="body">
				<ol>
					<li>
						<mark>weed</mark>
						<small>2767</small>
					</li>
					<li>
						<mark>LdDoxwOz1F</mark>
						<small>2719</small>
					</li>
					<li>
						<mark>Ys8</mark>
						<small>2658</small>
					</li>
					<li>
						<mark>lAIs88K4e3ZX</mark>
						<small>2654</small>
					</li>
					<li>
						<mark>BRkEpf3y</mark>
						<small>2651</small>
					</li>
					
				</ol>
			</div>
		</div>
	</div>
    )
}
export default Leaderboard;