export interface IFilm {
    UUID:       string
    Name:       string
    Release:    number
    Grade:      number
    Genre:      string
    Price:      number
    WhatchTime: number
    Summary:    string
    Image:      string
    Video:      string
}

export interface ICart {
    FilmUUID:       string
    Quantity:   number
}

export interface IOrder {
    UUID: string
    Film: string
    Quantity: number
    UserUUID: string
    Date: string
    Status: string
}
