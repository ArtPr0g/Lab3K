import {useContext} from "react";
import {GetUser} from "../requests/GetUser";
import {updateStatus} from "../modules";
import {ContextOrder} from "../requests/GetOrders";


export function Order() {
    const ctx = useContext(ContextOrder)

    const handleChangeStatus = (event: { target: { value: any; }; }) => {
        ctx.Status = event.target.value
        updateStatus(ctx.UUID, ctx.Status)
        window.location.replace('/orders')
    };

    return (
        <div
            className="border-2 border-teal-200 mx-auto mt-4 w-2/3 h-20 rounded-lg grid grid-cols-5 bg-white"
        >
            <p className="text-red-600 place-self-center text-2xl font-medium">
                {ctx.Film}
            </p>

            <p className="text-blue-700 place-self-center text-2xl font-medium">
                {ctx.Quantity}
            </p>

            <p className="text-green-500 place-self-center text-2xl font-medium">
                {GetUser(ctx.UserUUID)}
            </p>

            <p className="text-yellow-400 place-self-center text-1xl font-medium">
                {ctx.Date.replace("T", " ").split(".")[0]}
            </p>

            <div className="text-orange-500 place-self-center font-medium">
                <select
                    onChange={handleChangeStatus}
                    value={ctx.Status}
                    className="mt-1 w-40 rounded-md border border-gray-300 bg-white py-2 px-3 focus:border-gray-500 focus:outline-none focus:ring-gray-500"
                >
                    <option>Заказан</option>
                    <option>Оплачен</option>
                    <option>Подтверждён</option>
                    <option>Отдан</option>
                </select>
            </div>

        </div>
    )
}