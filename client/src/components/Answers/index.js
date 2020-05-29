import React, { useEffect, useState } from "react";
import moment from 'moment'
import "./style.css"

function Answers({ data }) {

    const [list, setList] = useState([])

    useEffect(() => {
        if (!data) return
        setList(data)
        list.map(item => (console.log(item)))
    }, [list, data])

    return (
        <>
            {list.map(item => (
                <ul key={item.Id} className="answer">
                    <strong>{item.Text}</strong>
                    <div>
                        <span className="like">likes {item.Likes}</span>
                        <span>{moment.unix(item.Date / 1000).utc().format('DD-MM-YYYY')}</span>
                    </div>
                </ul>
            ))}
        </>

    );
}

export default Answers;
