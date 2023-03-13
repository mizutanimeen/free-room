import { useState } from "react";

import React from "react";
import Button from "./components/Button";
import basePath from "./apis/basePath";

const App = () => {
    const [resources, setResources] = useState({"status":""});

    const get = async() => {
        try {
            const get = await basePath.get("");
            setResources(get.data)
        } catch (err) {
            console.log(err);
        }
    };

    return (
        <div className="ui container" style={{margin:"20px"}}>
            <Button onClick={get} color="primary" text="GET" />
            <div>{resources["status"]}</div>
        </div>
    );
};

export default App;