import {Link} from "react-router-dom";
import {ICart} from "../models";
import {Cart} from "./Cart";
import {GetCart, ContextCart} from "../requests/GetCart";
import React from "react";
import {Navbar} from "./Navbar";

export function CartPage() {
    return (
        <>
        <Navbar/>
        <div className="bg-yellow-50 min-h-screen">
            <p className="ml-4 sm:text-2xl text-1xl font-normal text-black">
                <Link to="/film" className="mr-2">
                    OnlineKino
                </Link>
                / cart
            </p>

            <p className="text-center sm:text-5xl text-3xl font-bold text-pink-500">
                Корзина
            </p>

            {GetCart().map((cart: ICart) => {
                return(
                    <ContextCart.Provider value={cart}>
                        <Cart/>
                    </ContextCart.Provider>
                )
            })}

            <p className="my-8 text-center">
                <Link to="/film"
                      className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                >
                    Обратно на главную
                </Link>
            </p>
        </div>
            </>
    )
}