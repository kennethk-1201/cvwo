import { useState, useEffect } from "react";

import Axios from "axios";

function App() {

    const [Tasks, setTasks] = useState([])

    useEffect(() =>{
        const initialiseApp = async () => {
            const res = await Axios.get(`${process.env.REACT_APP_BACKEND_API}/read`)
            console.log(res);
        }
        initialiseApp()
    }, [])

    return (
        <div className="App">
        </div>
    );
}

export default App;