import { useState, useEffect } from "react";

import Axios from "axios";

import {
    Loader,
    CreateTask,
    Footer,
    Header,
    TaskContainer
} from "./components";

function App() {

    const [tasks, setTasks] = useState([])
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getTasks = async () => {
            const res = await Axios.get(`${process.env.REACT_APP_BACKEND_API}/read`)
            console.log(res);
            setTasks(res.data.data)
            setLoading(false);
        }
        setTimeout(getTasks, 0);
    }, [])

    return (
        <div className="container">
            <Header/>
            {loading ? <Loader/> : <TaskContainer tasks={tasks}/>}
            <Footer/>
        </div>
    );
}

export default App;