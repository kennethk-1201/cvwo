import { useState } from "react";
import Axios from "axios";

import CloseIcon from '@mui/icons-material/Close';
import { sortTasks } from "../../helpers";

const ModalForm = props => {
    
    const {
        tasks, 
        setTasks, 
        setLoading,
        currentId,
        setCurrentId,
        title,
        setTitle,
        description,
        setDescription,
        deadline,
        setDeadline
    } = props

    const closeModal = e => {
        window.document.getElementById("modal").classList.add("d-none");

        // clear form state
        setCurrentId(-1);
        setTitle("");
        setDescription("");
        setDeadline("");
    }
    
    const submitHandler = async e => {
        e.preventDefault();
        try {
            // if invalid, do not send request
            if (title == "" || description == "" || deadline == "" || title.length > 100 || description.length > 500) {
                return;
            }
            closeModal();

            setLoading(true);
            
            if (currentId === -1){
                // if currentId does not exist, send POST to create task
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
            } else {
                // else, send PUT to update task
                const res = await Axios.put(
                    `${process.env.REACT_APP_BACKEND_API}/update/${currentId}`,
                    {
                        title,
                        body: description,
                        deadline
                    }
                )
                // update frontend
                setTasks(sortTasks([res.data.data, ...tasks.filter(t => t.id !== currentId)]));
            }

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
                <div>
                <input className="modal-form-deadline" type="datetime-local" value={deadline} onChange={e => changeHandler(e, setDeadline)} placeholder="Task title"/>
                </div>
                <div className="modal-form-button-holder">
                    <button className="modal-form-button" type="submit">Save</button>
                </div>
            </form>
        </div>
    )
}

export default ModalForm;