import {ENDPOINT} from "./App";
import axios from "axios";

export function getToken() {
    let tokens = document.cookie.split(' ')
    let access_token = ''
    for (var i = 0; i < tokens.length; i++) {
        if (tokens[i].startsWith("access_token=")) {
            access_token = tokens[i].replace("access_token=", "")
        }
    }
    return access_token.replace(";", "")
}

export function getRole() {
    let access_token = getToken()
    return axios.get(`${ENDPOINT}/role`, {
        withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }
    }).then(r => r.data)
}

export function getFromBackend(url: string) {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data)
}

export function getFromBackendToken(url: string) {
    let access_token = getToken()
    return axios.get(`${ENDPOINT}/${url}`, {
        withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }
    }).then(r => r.data)
}

export function deleteFromBackendToken(url: string) {
    let access_token = getToken()
    return axios.delete(`${ENDPOINT}/${url}`, {
        withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }
    }).then(r => r.data)
}

export function addFilm(url: string, name: string, release: number, grade: number, genre: string,  price: number, whatch_time: number, summary: string, video: string, image: string) {
    const body = {
        Name: name,
        Release: release,
        Grade: grade,
        Genre: genre,
        Price: price,
        WhatchTime: whatch_time,
        Summary: summary,
        Video: video,
        Image: image,
    }
    let access_token = getToken()
    console.log(body)
    return axios.post(`${ENDPOINT}/${url}`, body, {
        withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }
    }).then(function (response) {
        console.log(response);
    })
}

export function updateStatus(uuid: string, status: string) {
    const body = { Status: status }
    let access_token = getToken()
    return axios.put(`${ENDPOINT}/orders/${uuid}`, body,{withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(r => r.data)
}

export function deleteFilm (url: string, uuid: string) {
    let access_token = getToken()
    return axios.delete(`${ENDPOINT}/${url}/${uuid}`, {withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(r => r.data)
}

export function changeFilm(uuid: string, url: string, name: string, release: number, grade: number, genre: string,  price: number, whatch_time: number, summary: string, video: string, image: string) {
    const body = {
        Name: name,
        Release: release,
        Grade: grade,
        Genre: genre,
        Price: price,
        WhatchTime: whatch_time,
        Summary: summary,
        Video: video,
        Image: image,
    }
    let access_token = getToken()
    return  axios.put(`${ENDPOINT}/${url}/${uuid}`, body, {withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(function (response) {
        console.log(response);
    })

}

export function createUser(url: string, name: string, pass: string) {
    const body = {name: name, pass: pass}
    return axios.post(`${ENDPOINT}/${url}`, body, {withCredentials: true}).then(function (response) {
        console.log(response)
        window.location.replace("/login")
    }).catch(function () {
        window.location.replace("/registration")
    })
}

export function loginUser(url: string, name: string, pass: string) {
    const body = {login: name, password: pass}
    return axios.post(`${ENDPOINT}/${url}`, body, {withCredentials: true}).then(function (response) {
        console.log(response)
        window.location.replace("/film")
    }).catch(function () {
        window.location.replace("/login")
    })
}

export function logoutUser(url: string) {
    let access_token = getToken()
    return axios.get(`${ENDPOINT}/${url}`, {
        withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }
    }).then(function () {
        window.location.replace('/login')
    })
}
