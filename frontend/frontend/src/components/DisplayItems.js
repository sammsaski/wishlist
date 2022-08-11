import React, { useEffect, useState } from 'react'
import axios from 'axios';

export default function DisplayItems() {
    const [items, setItems] = useState([]);

    //
    // function buildTable() {
    //     // create table
    //     const table = <table></table>;
    //     table.appendChild(thead);
    //     table.appendChild(<tbody id="table-body"></tbody>)

    //     // create table body
    //     const tableBody = document.getElementById('table-body');
    //     const tableRow = <tr></tr>;
    //     items.forEach((item) => {
    //         // fill id
    //         tableRow.appendChild(<td>{item.id}</td>)

    //         // fill name
    //         tableRow.appendChild(<td>{item.name}</td>)

    //         // fill price
    //         tableRow.appendChild(<td>{item.price}</td>)
    //     });
    //     tableBody.appendChild(tableRow);

    //     return table;
    // }

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
