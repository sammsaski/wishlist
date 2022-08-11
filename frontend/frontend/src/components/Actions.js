import '../static/actions.css'
import { useState } from 'react';
import CreateItem from './CreateItem';
import DeleteItem from './DeleteItem';

export default function Actions() {
    const [isCreate, setIsCreate] = useState(true);

    const handleChange = event => {
        setIsCreate(!isCreate);
    }

    return (
        <div>
            <div>
                <input id="choice" type="radio" name="choice" checked={isCreate === true} onClick={handleChange}></input>
                <label>Create</label>

                <input id="choice" type="radio" name="choice" checked={isCreate === false} onClick={handleChange}></input>
                <label>Delete</label>
            </div>

            <div>
                {isCreate ? (
                    <CreateItem></CreateItem>
                ) : (
                    <DeleteItem></DeleteItem>
                )}
            </div>
        </div>
    )
}