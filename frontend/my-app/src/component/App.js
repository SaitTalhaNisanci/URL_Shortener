import React, { Component } from 'react';
import './App.css';
import axios from 'axios';

class App extends Component {

    constructor(props){
        super(props);

        this.state = {
            long_url : '',
            short_url : ''
        }

        this.updateInput = this.updateInput.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }


    updateInput(event){
        this.setState({long_url : event.target.value})
    }


    handleSubmit(){
        console.log('Your input value is: ' + this.state.long_url)
        // TODO :: make this call simpler
        axios({
            method: 'post',
            url: 'http://localhost:8080/shorten?url=' + this.state.long_url,
            headers: {
                'Content-Type': 'application/json',
            },
        }).then( resp => {
            this.setState({long_url: this.state.long_url, short_url:resp.data['Short']});
        }).catch( err => {
            console.log(err);
        });
    }



    render(){
        return (
            <div>
                <input type="text" onChange={this.updateInput}></input>
                <input type="submit" value="Short URL" onClick={this.handleSubmit} ></input>
                {this.state.short_url}
            </div>
        );
    }
}



export default App;
