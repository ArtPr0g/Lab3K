import React, {useState} from 'react';
import {LoginUser} from "../requests/LoginUser";


export function LoginPage() {
    const [name, setName] = useState('');

    const handleChangeName = (event: { target: { value: any; }; }) => {
        setName(event.target.value);

    };
    const [pass, setPass] = useState('');

    const handleChangePass = (event: { target: { value: any; }; }) => {
        setPass(event.target.value);
    };
    return (
        <div className="relative flex flex-col justify-center min-h-screen overflow-hidden">
            <div className="w-full p-6 m-auto bg-yellow-50 rounded-md shadow-md lg:max-w-xl">
                <h1 className="text-3xl font-semibold text-center text-black">
                    Войти
                </h1>
                <form className="mt-6">
                    <div className="mb-2">
                        <label
                            htmlFor="login"
                            className="block text-sm font-semibold text-black"
                        >
                            Логин
                        </label>
                        <input
                            type="login"
                            onChange={handleChangeName}
                            value={name}
                            className="block w-full px-4 py-2 mt-2 text-black bg-white border rounded-md focus:border-black focus:ring-black focus:outline-none focus:ring focus:ring-opacity-40"
                        />
                    </div>
                    <div className="mb-2">
                        <label
                            htmlFor="password"
                            className="block text-sm font-semibold text-black"
                        >
                            Пароль
                        </label>
                        <input
                            type="password"
                            onChange={handleChangePass}
                            value={pass}
                            className="block w-full px-4 py-2 mt-2 text-black bg-white border rounded-md focus:border-black focus:ring-black focus:outline-none focus:ring focus:ring-opacity-40"
                        />
                    </div>
                </form>
                <div className="mt-6">
                    {LoginUser(name, pass)}
                </div>

                <p className="mt-8 text-xl font-light text-center text-black">
                    {" "}
                    Отсутствует аккаунт?{" "}
                    <a
                        href="/registration"
                        className="font-medium text-black text-xl hover:underline"
                    >
                        Зарегистрироваться
                    </a>
                </p>
            </div>
        </div>
    );
}