import React, { Component } from 'react';
import './App.css';
import './styles.css';
import axios from 'axios';
import {Button, Alert, FormControl} from 'react-bootstrap';
import {CopyToClipboard} from 'react-copy-to-clipboard';
import serverIP from './constants';

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
            url: serverIP + ':8080/shorten?url=' + this.state.long_url,
            headers: {
                'Content-Type': 'application/json',
            },
        }).then( resp => {
            this.setState({long_url: this.state.long_url, short_url:serverIP + ':3000/' + resp.data['Short']});
        }).catch( err => {
            console.log(err);
        });
    }



    render(){
        return (
            <div>
                <div className="wrapper">
                    <div className="main">
                        <Alert bsStyle="warning" className="text-center">
                            <strong>Welcome!</strong> Enter URL to shorten it.
                        </Alert>
                        <FormControl
                            className="form"
                            type="text"
                            placeholder="Enter URL to shorten"
                            onChange={this.updateInput}
                        />

                        <div >
                            <Button className= "button" bsStyle="success" onClick={this.handleSubmit}>Shorten</Button>
                        </div>

                        <CopyToClipboard text={this.state.short_url}>
                            <div>
                                <Button class="button" bsStyle="primary" >Copy</Button>
                                <p className="text-center">{this.state.short_url}</p>
                            </div>

                        </CopyToClipboard>
                    </div>
                </div>
            </div>
        );
    }
}



export default App;
