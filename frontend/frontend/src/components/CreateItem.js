import React, { useState } from 'react'
import axios from 'axios';

export default function CreateItem() {
    const [name, setName] = useState('')
    const [price, setPrice] = useState(0)

    const handleSubmit = event => {
        event.preventDefault();
        const item = {
            name: name,
            price: price
        }

        axios.post('items', item)
            .then((response) => {
                console.log(response);
                console.log(response.data);
            })
        
        // clear the forms
        const inputs = document.getElementsByName('input');
        inputs.forEach((input) => {
            input.value = '';
        });
    }

    /**
     * Data type of name must be string.
     */
    const handleNameChange = event => {
        setName(event.target.value)
    }

    /**
     * Data type of price must be integer.
     */
    const handlePriceChange = event => {
        setPrice(parseInt(event.target.value))
    }

    return (
        <form onSubmit={handleSubmit}>
            <label>Name:
                <input type="text" name="input" onChange={handleNameChange} />
            </label>

            <label>Price:
                <input type="text" name="input" onChange={handlePriceChange}/>
            </label>

            <input type="submit" value="Submit"></input>
        </form>
    )
}