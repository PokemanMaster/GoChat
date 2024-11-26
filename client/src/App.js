import {useRoutes} from "react-router-dom"
import router from "../../../../../ECP电商商城项目/Golang-Project/client/src/router";
import React from 'react'


export default function App() {
    const tourist = useRoutes(router)
    return (
        <>
            {tourist}
        </>
    )
}


