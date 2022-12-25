import {loginUser} from "../modules";
import React from "react";


export function LoginUser(name: string, pass: string) {
    const url = `login`

    function Login() {
        loginUser(url, name, pass)
    }

    return (
        <button
            className="w-full text-center px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-black rounded-md hover:bg-white hover:text-black focus:outline-none focus:bg-white"
            onClick={() => Login()}
        >
            Войти
        </button>
    )
}