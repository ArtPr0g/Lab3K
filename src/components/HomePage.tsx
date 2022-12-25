import {Film} from "./Film"
import {GetFilms, ContextFilm} from "../requests/GetFilms";
import {IFilm} from "../models";
import React, {useState} from "react";
import {Box, Slider} from "@mui/material";
import {Navbar} from "./Navbar";

export function HomePage() {
    const films = GetFilms()

    const [name, setName] = useState('')

    const filteredFilms = films.filter((film: { Name: string }) => {
        return film.Name.toLowerCase().includes(name.toLowerCase())
    })

    const [price, setPrice] = React.useState<number[]>([0, 1000])

    const minDistance = 10;

    const handleChange = (event: Event, newValue: number | number[], activeThumb: number,) => {
        if (!Array.isArray(newValue)) {
            return
        }

        if (activeThumb === 0) {
            setPrice([Math.min(newValue[0], price[1] - minDistance), price[1]])
        } else {
            setPrice([price[0], Math.max(newValue[1], price[0] + minDistance)])
        }
    }

    const marks = [
        {
            value: 0,
            label: '0 ₽',
        },
        {
            value: 50,
            label: '50 ₽',
        },
        {
            value: 100,
            label: '100 ₽',
        },
        {
            value: 200,
            label: '200 ₽',
        },
        {
            value: 500,
            label: '500 ₽',
        },
    ];

    function valuetext(price: number) {
        return `${price} Р`;
    }

    return (
        <>
            <Navbar/>
        <div className="bg-yellow-50 min-h-screen">
            <p className="ml-4 sm:text-2xl text-1xl font-normal text-black">
                OnlineKino
            </p>

            <p className="text-center sm:text-5xl text-3xl font-bold text-pink-500">
                Фильмы
            </p>

            <div className="mt-5 mob:mt-2">
                <div className="flex place-content-center">
                    <Box sx={{width: 400}}>
                        <Slider
                            aria-label="Price filter"
                            valueLabelDisplay="auto"
                            getAriaValueText={valuetext}
                            value={price}
                            marks={marks}
                            onChange={handleChange}
                            disableSwap
                            step={10}
                            min={0}
                            max={500}
                        />
                    </Box>
                </div>

                <div className="flex place-content-center">
                    <form>
                        <input
                            type="text"
                            className="block w-full px-4 py-2 text-gray-500 text-2xl bg-white border rounded-full focus:border-gray-400 focus:ring-gray-400 focus:outline-none focus:ring focus:ring-opacity-40"
                            placeholder="Поиск..."
                            onChange={(event) => setName(event.target.value)}
                        />
                    </form>
                </div>
            </div>


            <div className="pt-5 flex flex-col gap-4  mx-auto container">
                {filteredFilms.filter((film: { Price: number; }) => film.Price >= price[0] && film.Price <= price[1]).map((film: IFilm) => {
                    return (
                        <ContextFilm.Provider value={film}>
                            <Film/>
                        </ContextFilm.Provider>
                    )
                })}

            </div>
        </div>
            </>
    )
}