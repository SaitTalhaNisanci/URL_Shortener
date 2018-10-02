import React, { Component } from 'react';
import axios from "axios/index";
import serverIP from './constants';

class Original extends Component {
    constructor(props){
        super(props);

        this.state = {
            long_url : '',
            short_url : ''
        }
        var request_url =serverIP + ':8080/original?short_url=' + this.props.match.params.id;
        axios.get(request_url)
            .then(res => {
                window.location.assign(res.data['Long']);
            })
            .catch(err => {
                console.log(err);
            })
    }

    render(){
        return (
            <div>
                <input type="text"></input>
                {this.state.short_url}
            </div>
        );
    }
}


export default Original;