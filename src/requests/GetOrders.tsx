import {createContext} from "react";
import {OrderContext} from "../context";

export const ContextOrder = createContext(OrderContext);
const success = "Success"

export function reducer(state: any, action: { type: any; payload: any; }) {
    switch (action.type) {
        case success:
            return {
                order: action.payload
            }
        default:
            return state
    }
}