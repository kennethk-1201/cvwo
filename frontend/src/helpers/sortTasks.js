const sortTasks = tasks => {
    return tasks.sort(function (a, b) {
        return (a.deadline > b.deadline) - (a.deadline < b.deadline);
    });
}

export default sortTasks;