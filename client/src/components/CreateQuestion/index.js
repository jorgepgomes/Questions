import React, { useState } from "react";
import "./style.css"

function CreateQuestion({ onSubmit }) {
    const [text, setText] = useState("");
    const [user, setUser] = useState("");

    async function handleSubmit(e) {
        e.preventDefault();
        try {
            let res = await onSubmit({ text, user })
            console.log("create user: ", res)

        } catch (err) {
            console.log(err)
        }
    }

    return (
        <>
            <strong>Fazer uma Pergunta</strong>
            <form onSubmit={handleSubmit}>
                <div className="input-block">
                    <label>User</label>
                    <input
                        name="use"
                        value={user}
                        type="text"
                        onChange={e => setUser(e.target.value)}
                    />
                </div>
                <div className="input-block">
                    <label>Pergunta</label>
                    <input
                        name="text"
                        value={text}
                        type="text"
                        onChange={e => setText(e.target.value)}
                    />
                </div>
                <button type="submit">Criar</button>
            </form>
        </>
    );
}

export default CreateQuestion;
