import React, { useState } from "react";
import "./style.css"

function Search({ onSubmit }) {
    const [search, setSearch] = useState("");

    async function handleSubmit(e) {
        e.preventDefault();
        await onSubmit(search);
    }
    return (
        <div className="form-search">
            <form onSubmit={handleSubmit}>
                <div className="">
                    <input
                        name="search"
                        value={search}
                        type="text"
                        onChange={e => setSearch(e.target.value)}
                    />
                    <button type="submit">Pesquisar</button>

                </div>
            </form>
        </div>
    );
}

export default Search;
