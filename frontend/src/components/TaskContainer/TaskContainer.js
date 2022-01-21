import { Task, AddButton } from "..";

const TaskContainer = props => {

    const { tasks, setTasks } = props;

    return (
        <div className="tasklist-container fade-in">
            <AddButton/>
            <div className="tasks-holder">
                {
                    tasks.map((t, i) => <Task key={i} task={t} tasks={tasks} setTasks={setTasks}/>)
                }
            </div>
        </div>
    )
}

export default TaskContainer;