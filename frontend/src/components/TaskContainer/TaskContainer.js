import { Task, AddButton } from "..";

const TaskContainer = props => {

    const { tasks, setTasks, setLoading } = props;

    return (
        <div className="tasklist-container fade-in">
            <AddButton/>
            <div className="tasks-holder">
                {
                    tasks.map((t, i) => <Task key={i} task={t} tasks={tasks} setTasks={setTasks} setLoading={setLoading}/>)
                }
            </div>
        </div>
    )
}

export default TaskContainer;