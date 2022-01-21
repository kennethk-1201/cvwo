import Axios from "axios";

import EditIcon from '@mui/icons-material/Edit';
import DoneIcon from '@mui/icons-material/Done';

import { formatTime } from "../../helpers";

const Task = props => {

    const { 
        task, 
        tasks, 
        setTasks, 
        setTitle,
        setDescription,
        setDeadline,
        setCurrentId,
        setLoading 
    } = props;

    const deleteTask = async e => {
                
        setLoading(true);

        const res = await Axios.delete(`${process.env.REACT_APP_BACKEND_API}/delete/${task.id}`, {
            headers: {}, data: {}
        })

        // update frontend
        setTasks(tasks.filter(t => t.id !== task.id))

        setLoading(false);
    }

    const openModal = e => {
        setCurrentId(task.id);
        setTitle(task.title);
        setDescription(task.body);
        setDeadline(task.deadline.slice(0,16));
        window.document.getElementById("modal").classList.remove("d-none");
    }

    return (
        <div className="task-container">
            <h3 className="task-title">{task.title}</h3>
            <p className="task-body">{task.body}</p>
            <div className="task-bottom-row">
                <p className="task-deadline">{formatTime(task.deadline)}</p>
                <div className="icon-row">
                    <EditIcon color="black" className="edit-icon" onClick={openModal}/>
                    <DoneIcon color="black" className="done-icon" onClick={deleteTask}/>
                </div>
            </div>
        </div>
    )
}

export default Task;