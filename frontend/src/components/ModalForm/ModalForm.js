import { useState } from "react";
import Axios from "axios";

import CloseIcon from '@mui/icons-material/Close';
import { sortTasks } from "../../helpers";

const ModalForm = props => {
    
    const {tasks, setTasks, setLoading} = props

    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [deadline, setDeadline] = useState("");

    const closeModal = e => {
        window.document.getElementById("modal").classList.add("d-none");
    }
    
    const submitHandler = async e => {
        e.preventDefault();
        try {
            closeModal();

            // POST request
            setLoading(true);
            const res = await Axios.post(
                `${process.env.REACT_APP_BACKEND_API}/create`,
                {
                    title,
                    body: description,
                    deadline
                }
            )

            // update frontend
            setTasks(sortTasks([res.data.data, ...tasks]));
            setLoading(false);
            
        } catch (err) {
            console.log("Error: ", err);
        }
    }

    const changeHandler = (e, setFn) => {
        setFn(e.target.value);
    }

    return (
        <div className="modal d-none fade-in" id="modal">
            <div className="modal-container-bg" onClick={closeModal}>
                
            </div>
            <form className="modal-container" onSubmit={submitHandler}>
                <div className="close-icon-holder">
                    <CloseIcon onClick={closeModal}/>
                </div>
                <input className="modal-form-title" type="text" value={title} onChange={e => changeHandler(e, setTitle)}  placeholder="Task title"/>
                <input className="modal-form-body" type="text" value={description} onChange={e => changeHandler(e, setDescription)} placeholder="Description"/>
                <input className="modal-form-deadline" type="datetime-local" value={deadline} onChange={e => changeHandler(e, setDeadline)} placeholder="Task title"/>
                <div className="modal-form-button-holder">
                    <button className="modal-form-button" type="submit">Save</button>
                </div>
            </form>
        </div>
    )
}

export default ModalForm;