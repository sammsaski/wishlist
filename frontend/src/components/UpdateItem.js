import React, { useState } from 'react'
import axios from 'axios';

export default function UpdateItem() {
    const [itemID, setItemID] = useState(0);
    const [name, setName] = useState('');
    const [price, setPrice] = useState(0);

    /**
     * Submit request to delete item with ID = itemID.
     */
     const handleSubmit = event => {
        event.preventDefault();

        const itemToUpdate = `items/${itemID}`
        const item = {
            name: name,
            price: price
        }

        axios.put(itemToUpdate, item)
            .then((response) => {
                console.log(response)
            });
        
        // clear the forms
        const inputs = document.getElementsByName('input');
        inputs.forEach((input) => {
            input.value = '';
        });
    }

    /**
     * Update state as form gets updated.
     */
    const handleChange = event => {
        setItemID(parseInt(event.target.value));
    }

    /**
     * Data type of name must be string.
     */
     const handleNameChange = event => {
        setName(event.target.value);
    }

    /**
     * Data type of price must be integer.
     */
    const handlePriceChange = event => {
        setPrice(parseInt(event.target.value));
    }

    return (
        <form onSubmit={handleSubmit}>
            <label>ID of Product to Update:
                <input type="text" onChange={handleChange}></input>
            </label>
            
            <label>New Name:
                <input type="text" onChange={handleNameChange}></input>
            </label>

            <label>New Price:
                <input type="text" onChange={handlePriceChange}></input>
            </label>

            <input type="submit" value="Submit"></input>
        </form>
    )
}