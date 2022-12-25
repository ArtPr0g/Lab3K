import {createContext, useEffect, useReducer} from "react";
import {getFromBackend} from "../modules";
import {FilmContext} from "../context";


export const ContextFilm = createContext(FilmContext);
const initialState = {films: []}
const success = "Success"

function reducer(state: any, action: { type: any; films: any; }) {
    switch (action.type) {
        case success:
            return {
                films: action.films
            }
        default:
            return state
    }
}

export function GetFilms() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `film`

    useEffect(() => {
        getFromBackend(url).then((result) => {
            dispatch({type: success, films: result})
        })
    }, [url])

    return state.films
}