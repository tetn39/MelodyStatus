import { useEffect, useState } from "react";
import axios from "axios";
import "./App.css";
import { Button } from "@/components/ui/button";

function App() {
	// usestate
	const [message, setMessage] = useState("");
	// localhost:8080/api/v1/helloからデータを取得する axiosを使う
	// 			setMessage(response.data.message); を追加
	useEffect(() => {
		axios.get("http://localhost:8080/api/v1/hello").then((response) => {
			console.log(response);
			setMessage(response.data.message);
		});
	}, []);

	return (
		<div className="App">
			<header className="App-header">
				<p>MelodyStatus v2</p>
				<a
					className="App-link"
					href="https://reactjs.org"
					target="_blank"
					rel="noopener noreferrer"
				>
					Learn React
				</a>
			<div className="text-5xl">{message}</div>
			<Button>Click me</Button>
			</header>
		</div>
	);
}

export default App;
