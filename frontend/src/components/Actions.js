import '../static/actions.css'
import { useState } from 'react';
import CreateItem from './CreateItem';
import DeleteItem from './DeleteItem';
import UpdateItem from './UpdateItem';


export default function Actions() {
    const [action, setAction] = useState('create');

    /**
     * Handle updating the current action that the user wants to perform.
     * @param {*} event 
     */
    const handleChange = event => {
        setAction(event.target.name);
    }

    /**
     * Switch between create/delete/update action menus.
     * @param {*} param : a string detailing which action the user would like to complete.
     * @returns the action menu corresponding to the desired param.
     */
    const renderSwitch = param => {
        switch(param) {
            case 'delete':
                return <DeleteItem></DeleteItem>;
            case 'update':
                return <UpdateItem></UpdateItem>;
            default:
                return <CreateItem></CreateItem>;
        }
    }

    return (
        <div>
            <div>
                <input className="choice" type="radio" name="create" checked={action === 'create'} onClick={handleChange}></input>
                <label>Create</label>

                <input className="choice" type="radio" name="delete" checked={action === 'delete'} onClick={handleChange}></input>
                <label>Delete</label>
                
                <input className="choice" type="radio" name="update" checked={action === 'update'} onClick={handleChange}></input>
                <label>Update</label>
            </div>

            <div>
                {renderSwitch(action)}
            </div>
        </div>
    )
}