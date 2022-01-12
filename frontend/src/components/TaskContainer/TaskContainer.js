import { Task } from "..";

const TaskContainer = props => {

    const { tasks } = props;

    return (
        <div className="tasklist-container fade-in">
            {
                tasks.map((t, i) => <Task key={i} task={t}/>)
            }
        </div>
    )
}

export default TaskContainer;