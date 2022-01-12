import { Task } from "..";

const TaskContainer = props => {

    const { tasks, setTasks } = props;

    return (
        <div className="tasklist-container fade-in">
            {
                tasks.map((t, i) => <Task key={i} task={t} tasks={tasks} setTasks={setTasks}/>)
            }
        </div>
    )
}

export default TaskContainer;