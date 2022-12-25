import React, {useState} from "react"
import {Navbar} from "./Navbar";
import {AddingFilm} from "../requests/AddFilm";
import {Link} from "react-router-dom";

export function AddFilm() {
    const [name, setName] = useState('');
    const handleChangeName = (event: { target: { value: any; }; }) => {
        setName(event.target.value);
    };

    const [release, setRelease] = useState(0);
    const handleChangeRelease = (event: { target: { value: any; }; }) => {
        setRelease(Number(event.target.value));
    };

    const [grade, setGrade] = useState(0);
    const handleChangeGrade = (event: { target: { value: any; }; }) => {
        setGrade(Number(event.target.value));
    };

    const [genre, setGenre] = useState('');
    const handleChangeGenre = (event: { target: { value: any; }; }) => {
        setGenre(event.target.value);
    };

    const [price, setPrice] = useState(0);
    const handleChangePrice = (event: { target: { value: any; }; }) => {
        setPrice(Number(event.target.value));
    };

    const [whatch_time, setWhatchTime] = useState(0);
    const handleChangeWhatchTime = (event: { target: { value: any; }; }) => {
        setWhatchTime(Number(event.target.value));
    };

    const [summary, setSummary] = useState("");
    const handleChangeSummary = (event: { target: { value: any; }; }) => {
        setSummary(event.target.value);
    };


    const [video, setVideo] = useState('');
    const handleChangeVideo = (event: { target: { value: string; }; }) => {
        setVideo(event.target.value);
    };

    const [image, setImage] = useState('');
    const handleChangeImage = (event: { target: { value: any; }; }) => {
        setImage(event.target.value);
    };

    return (
        <>
            <Navbar/>

            <div className="bg-yellow-50 min-h-screen">
                <p className="ml-4 sm:text-2xl text-1xl font-normal text-black">
                    <Link to="/film" className="mr-2">
                        OnlineKino
                    </Link>
                    / adding
                </p>

                <p className="text-center sm:text-5xl text-3xl font-bold text-pink-500">
                    Добавление флиьма
                </p>

                <form className="mt-10 mx-5 bg-white rounded-lg border-2 border-teal-200">
                    <div className="grid grid-cols-4 grid-rows-2 gap-10 p-8">
                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Название
                            </label>
                            <input
                                type="text"
                                onChange={handleChangeName}
                                value={name}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Год выхода фильма
                            </label>
                            <input
                                type="number"
                                min="1"
                                max="3000"
                                onChange={handleChangeRelease}
                                value={release}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Оценка
                            </label>
                            <input
                                type="number"
                                min="1"
                                max="10"
                                onChange={handleChangeGrade}
                                value={grade}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Жанр
                            </label>
                            <input
                                type="text"
                                onChange={handleChangeGenre}
                                value={genre}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>


                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Цена
                            </label>
                            <input
                                type="number"
                                min="10"
                                max="1000"
                                onChange={handleChangePrice}
                                value={price}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Длительность фильма(мин.)
                            </label>
                            <input
                                type="number"
                                min="1"
                                max="999999"
                                onChange={handleChangeWhatchTime}
                                value={whatch_time}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Краткое содержание
                            </label>
                            <input
                                type="text"
                                onChange={handleChangeSummary}
                                value={summary}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>


                        <div className="col-span-2">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Видео
                            </label>

                            <input
                                type="text"
                                onChange={handleChangeVideo}
                                value={video}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="col-span-2">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Изображение
                            </label>
                            <input
                                type="text"
                                onChange={handleChangeImage}
                                value={image}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>
                    </div>

                    <div className="text-center mb-6">
                        {AddingFilm(name, release, grade, genre,  price, whatch_time, summary, video, image)}
                    </div>
                </form>

                <p className="py-8 text-center">
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