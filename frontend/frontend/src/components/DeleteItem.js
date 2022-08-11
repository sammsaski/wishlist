import { useState } from 'react';
import axios from 'axios';

export default function DeleteItem() {
    const [itemID, setItemID] = useState(0)

    /**
     * Submit request to delete item with ID = itemID.
     */
    const handleSubmit = event => {
        const itemToDelete = `items/${itemID}`
        axios.delete(itemToDelete)
            .then((response) => {
                console.log(response)
            });
    }

    /**
     * Update state as form gets updated.
     */
    const handleChange = event => {
        setItemID(parseInt(event.target.value));
    }

    return (
        <form onSubmit={handleSubmit}>
            <label>ID of Product to Delete:
                <input type="text" onChange={handleChange}></input>
            </label>

            <input type="submit" value="Submit"></input>
        </form>
    )
}