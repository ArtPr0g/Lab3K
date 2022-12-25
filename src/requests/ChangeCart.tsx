import {useReducer} from "react";
import {deleteFromBackendToken, getFromBackendToken} from "../modules";
import DeleteIcon from '@mui/icons-material/Delete';

const increase = "Increase"
const del = "Delete"
const failure = "Failure"

function reducer(state: any, action: { type: any; payload?: any; }) {
    switch (action.type) {
        case increase:
            return {
                count: action.payload
            }
        case del:
            return {
                count: 0
            }
        case failure:
            return {
                count: 0
            }
        default:
            return state
    }
}


export function ChangeCart(Film: string) {
    const [dispatch] = useReducer(reducer, {count: 0});
    const url1 = `cart/increase/${Film}`
    const url2 = `cart/delete/${Film}`

    function Incr() {
        getFromBackendToken(url1).then(result => {
            dispatch({type: increase, payload: result})
        }).catch(() => {
            dispatch({type: failure})
        })
    }

    function Del() {
        deleteFromBackendToken(url2).then(() => {
            dispatch({type: del})
        }).catch(() => {
            dispatch({type: failure})
        })
    }

    return (
        <>
            <form className="inline">
                <button onClick={() => Incr()}>FAV</button>
                {" "}
                <button onClick={() => Del()}>
                    <DeleteIcon fontSize="inherit"/>
                </button>
            </form>
        </>
    );
}
