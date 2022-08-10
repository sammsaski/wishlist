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

    const data = items.map(item => {
        return <div>Name: {item.name}, Price: {item.price}</div>
    })

    return (
        <div>Data from API!
            { data }
        </div>
    )
}



// class DisplayItems extends Component {

//     constructor() {
//         super();
//         this.state = {"items": []};
//     }

    // componentWillMount() {
    //     axios.get('api/ping')
    //         .then((response) => {
    //             this.setState(() => {
    //                 return { pong: response.data.message }
    //             })
    //         })
    //         .catch(function (error) {
    //             console.log(error);
    //         });
    // }
    // componentWillMount() {
    //     axios.get('items')
    //         .then((response) => {
    //             this.setState(response.data["items"]);
    //         })
    //         .catch(function (error) {
    //             console.log(error);
    //         });
    // }

    // render() {
    //     return <h1>Hello World</h1>;
    // }
// }

// export default DisplayItems;