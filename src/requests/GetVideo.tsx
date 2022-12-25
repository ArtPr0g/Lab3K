import {useEffect, useReducer} from "react";
import {useLocation} from "react-router-dom";
import {getFromBackendToken} from "../modules";

const initialState = {video: ""}
const success = "Success"
const failure = "Failure"

function reducer(state: any, action: { type: any; payload?: any; }) {
    switch (action.type) {
        case success:
            return {
                video: action.payload
            }
        case failure:
            return {
                video: "ВСЁ!"
            }
        default:
            return state
    }
}

export function GetVideo() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `film/${useLocation().state.Film}/${useLocation().state.Quantity}`

    useEffect(() => {
        getFromBackendToken(url).then(result => {
            dispatch({type: success, payload: result})
        }).catch(() => {
            dispatch({type: failure})
        })
    }, [url])

    // return (
    //     <video controls width="100%">
    //         <source src={state.video} type="video/mp4" />
    //         Sorry, your browser doesn't support embedded videos.
    //     </video>
    // )
    return (
        <iframe
            width="560"
            height="315"
            src={state.video}
            title="Youtube Player"
            //frameborder="0"
            allowFullScreen
        />
    )
}