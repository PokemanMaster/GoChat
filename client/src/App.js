import {useRoutes} from "react-router-dom"

import React from 'react'
import router from "./router";


export default function App() {
    const tourist = useRoutes(router)
    return (
        <>
            {tourist}
        </>
    )
}


