import {ICart, IFilm, IOrder} from "./models";

export let FilmContext: IFilm = {
    UUID: "",
    Name: "",
    Release: 0,
    Grade: 0,
    Genre: "",
    Price: 0,
    WhatchTime: 0,
    Summary: "",
    Image: "",
    Video: "",
}

export let CartContext: ICart = {
    FilmUUID: "",
    Quantity: 0
}

export let OrderContext: IOrder = {
    UUID: "",
    Film: "",
    Quantity: 0,
    UserUUID: "",
    Date: "",
    Status: "",
}