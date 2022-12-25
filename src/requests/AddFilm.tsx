import {addFilm} from "../modules";
import React from "react";


export function AddingFilm(name: string, release: number, grade: number, genre: string,  price: number, whatch_time: number, summary: string, video: string, image: string) {

    const url = `film`

    function Add() {
        addFilm(url, name, release, grade, genre,  price, whatch_time, summary, video, image)
    }


    return (
        <>
            <button
                onClick={() => Add()}
                className="border-4 border-red-500 bg-white text-red-500 hover:bg-red-500 hover:text-white py-1 px-2 place-self-center rounded-full text-2xl font-bold"
            >
                Добавить
            </button>
        </>
    );

}