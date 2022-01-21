import AddIcon from '@mui/icons-material/Add';

const AddButton = props => {
    const clickHandler = e => {
        window.document.getElementById("modal").classList.remove("d-none");
    }
    return <div className="add-button-container">
        <div className="add-button" onClick={clickHandler}>
            <AddIcon/>
        </div>
    </div>
}

export default AddButton;