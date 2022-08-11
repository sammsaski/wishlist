import React, { useEffect, useState } from 'react'
import axios from 'axios';

export default function DisplayItems() {
    const [items, setItems] = useState([]);

    useEffect(() => {
        axios.get('items')
            .then((response) => {
                setItems(response.data["items"]);
            })
            .catch(function (error) {
                console.log(error)
            });
    })

    return (
        <div className="data-table">
            <table className="table">
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Price</th>
                </tr>
                {items.map((val, key) => {
                    return (
                        <tr key={key}>
                            <td>{val.id}</td>
                            <td>{val.name}</td>
                            <td>{val.price}</td>
                        </tr>
                    )
                })}
            </table>
        </div>
    )
}
