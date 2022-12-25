import {useEffect, useReducer} from "react";
import {getFromBackend} from "../modules";

const initialState = {film: ""}
const success = "Success"

function reducer(state: any, action: { type: any; film: any; }) {
    switch (action.type) {
        case success:
            return {
                film: action.film
            }
        default:
            return state
    }
}

export function GetFilm(uuid: string) {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `film/${uuid}`

    useEffect(() => {
        getFromBackend(url).then((result) => {
            dispatch({type: success, film: result})
        })
    }, [url])

    return state.film

}