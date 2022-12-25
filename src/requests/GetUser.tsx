import {useEffect, useReducer} from "react";
import {getFromBackendToken} from "../modules";

const initialState = {user: ""}
const success = "Success"

function reducer(state: any, action: { type: any; payload: any; }) {
    switch (action.type) {
        case success:
            return {
                user: action.payload
            }
        default:
            return state
    }
}

export function GetUser(uuid: string) {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `user/${uuid}`

    useEffect(() => {
        getFromBackendToken(url).then((result) => {
            dispatch({type: success, payload: result})
        })
    }, [url])

    return state.user
}