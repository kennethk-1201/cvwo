import { useState, useEffect } from "react";

import Axios from "axios";

import {
    Loader,
    Footer,
    Header,
    TaskContainer,
    ModalForm,
    Error
} from "./components";

import {
    sortTasks
} from "./helpers";

function App() {

    const [tasks, setTasks] = useState([])
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(false);

    useEffect(() => {
        const getTasks = async () => {
            try {
                const res = await Axios.get(`${process.env.REACT_APP_BACKEND_API}/read`)
                console.log(res);
                setTasks(sortTasks(res.data.data));
                setLoading(false);
            } catch (err) {
                setError(true);
                setLoading(false);
            }
        }
        setTimeout(getTasks, 0);
    }, [])

    return (
        <div className="container fade-in">
            <Header/>
            <ModalForm tasks={tasks} setTasks={setTasks} setLoading={setLoading}/>
            {loading 
                ? <Loader/> 
                : error 
                    ? <Error/>
                    : <TaskContainer tasks={tasks} setTasks={setTasks} setLoading={setLoading}/>}
            <Footer/>
        </div>
    );
}

export default App;