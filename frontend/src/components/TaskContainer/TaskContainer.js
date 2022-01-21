import { Task, AddButton } from "..";

const TaskContainer = props => {

    const { 
        tasks, 
        setTasks, 
        setTitle,
        setDescription,
        setDeadline,
        setCurrentId,
        setLoading
    } = props;

    return (
        <div className="tasklist-container fade-in">
            <AddButton/>
            <div className="tasks-holder">
                {
                    tasks.map((t, i) => <Task 
                        key={i} 
                        task={t} 
                        tasks={tasks} 
                        setTasks={setTasks} 
                        setTitle={setTitle}
                        setDescription={setDescription}
                        setDeadline={setDeadline} 
                        setCurrentId={setCurrentId}
                        setLoading={setLoading}
                    />)
                }
            </div>
        </div>
    )
}

export default TaskContainer;