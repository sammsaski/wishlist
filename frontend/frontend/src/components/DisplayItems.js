import React, { Component } from 'react'
import axios from 'axios';

class DisplayItems extends Component {

    constructor() {
        super();
        this.state = {
            pong: 'pending'
        }
    }

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
    componentWillMount() {
        axios.get('items')
            .then((response) => {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            });
    }

    render() {
        return <h1>Hello World</h1>;
    }
}

export default DisplayItems;