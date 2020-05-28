import React, { useState } from "react";

function Search({ onSubmit }) {
    const [search, setSearch] = useState("");

    async function handleSubmit(e) {
        e.preventDefault();
        await onSubmit(search);
    }
    return (
        <form onSubmit={handleSubmit}>
            <div className="input-block">
                <input
                    name="search"
                    value={search}
                    type="text"
                    onChange={e => setSearch(e.target.value)}
                />
                <button type="submit">Pesquisar</button>
            </div>
        </form>
    );
}

export default Search;
