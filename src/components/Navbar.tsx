import {LoginNavbar} from "./LoginNavbar";
import {GuestNavbar} from "./GuestNavbar";
import {getRole, getToken} from "../modules";
import {ManagerNavbar} from "./ManagerNavbar";
import {useState} from "react";


export function Navbar() {
    let access_token = getToken()
    const [roles, setRole] = useState()
    let role = getRole()
    role.then((result) => {
        setRole(result)
    })
    console.log()
    if (access_token === "") {
        return <GuestNavbar/>;
    } else if (roles === 1){
        return <ManagerNavbar/>;
    } else {
        return <LoginNavbar/>;
    }
}