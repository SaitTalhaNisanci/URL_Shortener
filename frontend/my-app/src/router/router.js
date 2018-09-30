import React from "react";
import { BrowserRouter, Route } from "react-router-dom";
import App from "../component/App";
import Original from "../component/original";
const Router = () => (
    <BrowserRouter>
        <div>
            <Route exact path="/" component={App} />
            <Route path="/:id" component={Original} />
        </div>
    </BrowserRouter>
);


export default Router;