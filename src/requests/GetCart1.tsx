import {useEffect, useReducer} from "react";
import {getFromBackendToken} from "../modules";

const initialState = {cart: ""}
const success = "Success"

function reducer(state: any, action: { type: any; cart: any; }) {
    switch (action.type) {
        case success:
            return {
                cart: action.cart
            }
        default:
            return state
    }
}

export function GetCart1(film: string) {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `cart/${film}`

    useEffect(() => {
        getFromBackendToken(url).then((result) => {
            dispatch({type: success, cart: result})
        })
    }, [url])

    return state.cart
}