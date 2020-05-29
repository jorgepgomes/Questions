import React, { useEffect, useState } from "react";
import moment from 'moment'

function ListQuestion({ data }) {

    const [list, setList] = useState([])

    useEffect(() => {
        onlyQuestions(data)
    }, [data])

    function onlyQuestions(data) {
        setList([])
        for (let n in data) {
            let object = {
                id: data[n].Id,
                text: data[n].Text,
                likes: data[n].Likes,
                date: data[n].Date,
                answers: data[n].Answers.length
            }

            setList(list => [...list, object])
        }
    }

    function nav(id) {
        window.location.href = `/${id}`;
    }

    return (
        <>
            {list.map(item => (
                <ul key={item.id} className="" onClick={() => nav(item.id)}>
                    <div className="">
                        <strong>{item.text}</strong>
                        <div>
                            <span>data: {moment.unix(item.date / 1000).utc().format('DD-MM-YYYY')}</span>
                            <span>respostas: {item.answers}</span>
                            <span>likes: {item.likes}</span>
                        </div>
                    </div>
                </ul>
            ))}
        </>

    );
}

export default ListQuestion;
