import { useState, useEffect } from "react";

import Axios from "axios";

import {
    CreateTask,
    Footer,
    Header,
    TaskContainer
} from "./components";

function App() {

    const [Tasks, setTasks] = useState([])

    useEffect(() =>{
        const initialiseApp = async () => {
            const res = await Axios.get(`${process.env.REACT_APP_BACKEND_API}/read`)
            console.log(res);
            setTasks(res.data.data)
        }
        initialiseApp()
    }, [])

    return (
        <div className="container">
            <Header/>
            <TaskContainer tasks={Tasks}/>
            <Footer/>
        </div>
    );
}

export default App;