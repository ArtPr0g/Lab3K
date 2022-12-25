import {Link} from "react-router-dom"
import {useLocation} from "react-router-dom"
import React from 'react';
import {GetVideo} from "../requests/GetVideo";
import {Navbar} from "./Navbar";

export function Payment() {
    return (
        <>
            <Navbar/>
        <div className="bg-yellow-50 min-h-screen">
            <p className="ml-4 sm:text-2xl text-1xl font-normal text-black">
                <Link to="/film" className="mr-2">
                    OnlineKino
                </Link>
                <Link to="/cart" className="mr-2">
                    / cart
                </Link>
                / {useLocation().state.Name}
            </p>

            <p className="text-center sm:text-5xl text-3xl font-bold text-pink-500">
                Страница оплаты
            </p>

            <p className="sm:mt-8 font-medium sm:text-4xl text-2xl text-green-500 text-center">
                Ваш фильм "{useLocation().state.Name}":
                <p className="font-bold italic sm:text-4xl text-2xl text-red-700">
                    {GetVideo()}
                </p>
            </p>

            <p className="py-8 text-center">
                <Link to="/cart"
                      className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                >
                    Обратно в корзину
                </Link>
            </p>

            <img src="https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/Payment_iw36pk.png"
                 width="23%" className="mx-auto" alt="Payment"/>
        </div>
            </>
    )
}