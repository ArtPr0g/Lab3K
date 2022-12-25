import {IOrder} from "../models";
import React, {useEffect, useReducer, useState} from "react";
import {Navbar} from "./Navbar";
import {Order} from "./Order";
import {getToken} from "../modules";
import axios from "axios";
import {ENDPOINT} from "../App";
import {ContextOrder, reducer} from "../requests/GetOrders";
import {Link} from "react-router-dom";


const initialState = {order: []}
const success = "Success"

export function OrderPage() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `orders`
    let access_token = getToken()

    const [filter, setFilter] = useState(false);
    const [stDate, setStDate] = useState('');
    const handleChangeStDate = (event: { target: { value: any; }; }) => {
        setStDate(event.target.value);
    };
    const [endDate, setEndDate] = useState('');
    const handleChangeEndDate = (event: { target: { value: any; }; }) => {
        setEndDate(event.target.value);
    };
    const [status, setStatus] = useState('Любой');
    const handleChangeStatus = (event: { target: { value: any; }; }) => {
        setStatus(event.target.value);
    };

    useEffect(() => {
        if (filter) {
            axios.get(`${ENDPOINT}/${url}`, {
                withCredentials: true, headers: {
                    "Authorization": `Bearer ${access_token}`
                }, params: {start_date: stDate, end_date: endDate, status: status}
            }).then(r => r.data).then((result) => {
                dispatch({type: success, payload: result})
            })
        } else {
            axios.get(`${ENDPOINT}/${url}`, {
                withCredentials: true, headers: {
                    "Authorization": `Bearer ${access_token}`
                }, params: {start_date: "", end_date: "", status: "Любой"}
            }).then(r => r.data).then((result) => {
                dispatch({type: success, payload: result})
            })
        }
    }, [filter])


    return (
        <>
            <Navbar/>

            <div className="bg-yellow-50 min-h-screen">
                <p className="ml-4 sm:text-2xl text-1xl font-normal text-black">
                    <Link to="/film" className="mr-2">
                        OnlineKino
                    </Link>
                    / order
                </p>

                <p className="text-center text-5xl mb-1 font-bold text-pink-500">
                    Заказы
                </p>

                <div className="w-3/5 mx-auto grid grid-rows-2 grid-cols-3 justify-items-center">
                    <div>
                        <label htmlFor="first-name" className="block text-lg text-center font-medium text-gray-700">
                            Статус
                        </label>

                        <select
                            onChange={handleChangeStatus}
                            value={status}
                            className="mt-1 block w-40 rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm focus:border-gray-500 focus:outline-none focus:ring-gray-500 sm:text-base"
                        >
                            <option>Любой</option>
                            <option>Заказан</option>
                            <option>Оплачен</option>
                            <option>Подтверждён</option>
                            <option>Отдан</option>
                        </select>
                    </div>

                    <div>
                        <label htmlFor="first-name" className="block text-lg text-center font-medium text-gray-700">
                            С даты
                        </label>
                        <input
                            type="text"
                            onChange={handleChangeStDate}
                            value={stDate}
                            maxLength={10}
                            className="mt-1 block w-40 rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm focus:border-gray-500 focus:outline-none focus:ring-gray-500 sm:text-base"
                        />
                    </div>

                    <div>
                        <label htmlFor="first-name" className="block text-lg text-center font-medium text-gray-700">
                            По дату
                        </label>
                        <input
                            type="text"
                            onChange={handleChangeEndDate}
                            value={endDate}
                            maxLength={10}
                            className="mt-1 block w-40 rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm focus:border-gray-500 focus:outline-none focus:ring-gray-500 sm:text-base"
                        />
                    </div>

                    <button
                        className="border-4 col-start-2 border-violet-500 text-violet-500 hover:bg-violet-500 hover:text-white py-1 px-3 place-self-center rounded-full text-2xl font-bold"

                        onClick={() => {
                            setFilter(!filter)
                        }}> {!filter && <p>Применить фильтр</p>}{filter && <p>Снять фильтр</p>}
                    </button>
                </div>

                <div className="pt-2 flex flex-col gap-0 mx-auto container">
                    <div className="border-2 border-gray-400 -mb-1 rounded-lg py-2  w-2/3 mx-auto grid grid-cols-5">
                        <p className="place-self-center text-2xl font-bold text-red-600">
                            Магазин
                        </p>

                        <p className="place-self-center text-2xl font-bold text-blue-700">
                            Количество
                        </p>

                        <p className="place-self-center text-2xl font-bold text-green-500">
                            Покупатель
                        </p>

                        <p className="place-self-center text-2xl font-bold text-yellow-400">
                            Дата
                        </p>

                        <p className="place-self-center text-2xl font-bold text-orange-500">
                            Статус
                        </p>
                    </div>
                    {state.order.map((order: IOrder, key: any) => {
                        return (
                            <ContextOrder.Provider value={order} key={key}>
                                <Order/>
                            </ContextOrder.Provider>
                        )
                    })}
                </div>
                <p className="py-8 text-center">
                    <Link to="/store"
                          className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                    >
                        Обратно на главную
                    </Link>
                </p>
            </div>
        </>
    )
}