const Task = props => {

    const { task } = props;
    
    return (
        <div className="task-container">
            <h3 className="task-title">{task.title}</h3>
            <p className="task-body">{task.body}</p>
            <p className="task-deadline">{task.deadline}</p>
        </div>
    )
}

export default Task;