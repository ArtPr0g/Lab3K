import {deleteFilm} from "../modules";


export function DeleteFilm(uuid: string) {

    const url = `film`

    function Delete() {
        deleteFilm(url, uuid)
    }


    return (
        <form>
            <button onClick={() => Delete()}>Удалить магазин</button>
        </form>
    );

}