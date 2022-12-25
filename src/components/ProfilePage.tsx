import {Navbar} from "./Navbar";
import React from "react";
import {Link} from "react-router-dom";

export function ProfilePage() {
    return (
        <>
            <Navbar/>
            <div className="bg-yellow-50 min-h-screen">
                <p className="ml-4 sm:text-2xl text-1xl font-normal text-black">
                    <Link to="/film" className="mr-2">
                        OnlineKino
                    </Link>
                    / profile
                </p>

                <p className="text-center sm:text-5xl text-3xl font-bold text-pink-500">
                    Профиль
                </p>

                <p className="mx-8 font-medium text-3xl text-indigo-700">
                    Это Ваш профиль
                </p>
            </div>
        </>
    )
}