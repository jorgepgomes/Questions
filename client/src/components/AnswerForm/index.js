import React, { useState } from "react";
import "./style.css"

function AnswerForm({ onSubmit }) {
    const [text, setText] = useState("");
    const [user, setUser] = useState("");

    async function handleSubmit(e) {
        e.preventDefault();
        try {
            await onSubmit({ text, user })            
        } catch (err) {
            console.log(err)
        }
    }

    return (
        <>
            <div className="form-answer">
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
                        <label>Resposta</label>
                        <input
                            name="text"
                            value={text}
                            type="text"
                            onChange={e => setText(e.target.value)}
                        />
                    </div>
                    <button type="submit">Responder</button>
                </form>

            </div>
        </>
    );
}

export default AnswerForm;
