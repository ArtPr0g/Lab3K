import {ChangeCart} from "../requests/ChangeCart";
import {GetFilm} from "../requests/GetFilm";
import {Link} from "react-router-dom";
import {useContext} from "react";
import {ContextCart} from "../requests/GetCart";

export function Cart() {
    const ctx = useContext(ContextCart)
    let Film = GetFilm(ctx.FilmUUID)

    return (
        <div
            className="border-2 border-teal-200 mx-auto mt-4 w-1/2 mob:w-11/12 h-40 mob:h-28 py-5 px-5 mob:py-2 mob:px-0 rounded-lg grid grid-rows-2 grid-cols-3 bg-white"
        >
            <img src={Film.Image}
                 className="place-self-center object-contain h-20 w-20 mob:h-12 mob:w-12" alt={Film.Name}
            />


            <p className="text-blue-700 place-self-center sm:text-2xl text-1xl font-bold mob:font-normal">
                Кол-во:{" "}{ctx.Quantity}{" "}{ChangeCart(ctx.FilmUUID)}
            </p>


            <p className="text-red-600 place-self-center sm:text-3xl text-1xl font-bold">
                {Film.Name}
            </p>

            <p className="text-yellow-400 place-self-center sm:text-2xl text-1xl font-bold">
                {Film.Price} ₽
            </p>

            <Link to={`/payment`}
                  className="border-4 mob:border-2 border-orange-500 text-orange-500 hover:bg-orange-500 hover:text-white py-1 px-3 place-self-center rounded-full sm:text-2xl text-1xl font-bold"
                  state={{Name: Film.Name, Film: ctx.FilmUUID, Quantity: ctx.Quantity}}
            >
                Купить
            </Link>
        </div>
    )
}