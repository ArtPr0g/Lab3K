import {useContext, useState} from "react";
import {ContextFilm} from "../requests/GetFilms";
import {GetCart1} from "../requests/GetCart1";
import {ChangeCart} from "../requests/ChangeCart";
import {getRole} from "../modules";
import {Link} from "react-router-dom";
import {DeleteFilm} from "../requests/DeleteFilm";

export function Film() {
    const ctx = useContext(ContextFilm)
    let Cart = GetCart1(ctx.UUID)
    const [roles, setRole] = useState()
    const role = getRole()
    role.then((result) => {
        setRole(result)
    })
    let showManagerButton = false
    if (roles === 1) {
        showManagerButton = true
    }
    return (
        <div
            className="border-2 border-teal-200 mx-auto mt-4 mob:mt-1 w-1/2 mob:w-11/12 h-56 mob:h-24 py-5 px-5 mob:py-2 mob:px-0 rounded-lg grid grid-rows-3 grid-cols-3 bg-white"
        >
            <img src={ctx.Image}
                 className="place-self-center object-contain h-20 w-20 mob:h-12 mob:w-12" alt={ctx.Name}
            />



            <p className="text-blue-700 place-self-center sm:text-2xl text-1xl font-bold mob:font-normal">
                В корзину:{" "}{Cart.Quantity}{" "}{ChangeCart(ctx.UUID)}
            </p>

            <p className="text-red-600 place-self-center  sm:text-3xl text-1xl font-bold">
                {ctx.Name}
            </p>


            <p className="text-yellow-400 place-self-center sm:text-2xl text-1xl font-bold">
                {ctx.Price} ₽
            </p>

            {showManagerButton && <Link to="/change"
                                        className="border-4 mob:border-2  border-violet-500 text-violet-500 hover:bg-violet-500 hover:text-white py-1 px-3 place-self-center rounded-full sm:text-2xl text1xl font-bold"
                                        state={{
                                            UUID:       ctx.UUID,
                                            Name:       ctx.Name,
                                            Release:    ctx.Release,
                                            Grade:      ctx.Grade,
                                            Genre:      ctx.Genre,
                                            Price:      ctx.Price,
                                            WhatchTime: ctx.WhatchTime,
                                            Summary:    ctx.Summary,
                                            Image:      ctx.Image,
                                            Video:      ctx.Video
                                        }}
            >
                Изменить
            </Link>}

            {showManagerButton &&
                <p className="border-4 mob:border-2 col-start-3 border-purple-500 text-purple-500 hover:bg-purple-500 hover:text-white py-1 px-3 place-self-center rounded-full sm:text-2xl text1xl font-bold">
                    {DeleteFilm(ctx.UUID)}
                </p>}
        </div>
    )
}