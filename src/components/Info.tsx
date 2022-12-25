import {Link} from "react-router-dom";
import {Navbar} from "./Navbar";
import React from "react";

export function Info() {
    return (
        <>
            <Navbar/>
        <div className="bg-yellow-50 min-h-screen">
            <p className="ml-4 sm:text-2xl text-1xl font-normal text-black">
                <Link to="/film" className="mr-2">
                    OnlineKino
                </Link>
                / info
            </p>

            <p className="text-center sm:text-5xl text-3xl font-bold text-pink-500">
                OnlineKino
            </p>

            <p className="text-center sm:mt-4 mx-8 font-medium mob:font-normal text-3xl mob:text-2xl text-indigo-700">
                Это лучшая платформа для простотра фильмов в хорошем качестве!
            </p>

            <p className="py-8 text-center">
                <Link to="/film"
                      className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                >
                    Обратно на главную
                </Link>
            </p>

            <img src="https://res.cloudinary.com/dbgomdrm8/image/upload/v1670253989/1618x1080_0x0a330c9f_3086277691515592251_j82a1c.jpg" width="29%" className="mx-auto" alt="Discount"/>
        </div>
            </>
    )
}