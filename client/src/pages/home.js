import React, { useEffect, useState } from 'react';
import server from '../services/server'
import Search from '../components/Search'

export default function Home() {
    const [list, setList] = useState([]);

    useEffect(() => {
        async function listQuestions() {
            try {
                let res = await server.ListQuestions("")
                setList(res.data)
            } catch (err) {
                console.log("err: ", err)
            }
        }
        listQuestions()
    }, [])

    async function searchListQuestions(search) {
        try {
            let res = await server.ListQuestions(search)
            setList(res.data)
        } catch (err) {
            console.log("err: ", err)
        }
    }

    return (
        <>
            <Search onSubmit={searchListQuestions} />
            <p>home</p>
            <p>{JSON.stringify(list)}</p>

        </>
    );
}
