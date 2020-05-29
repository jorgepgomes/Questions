import React, { useEffect, useState } from 'react';
import server from '../../services/server'
import Search from '../../components/Search'
import ListQuestion from '../../components/ListQuestion'
import CreateQuestion from '../../components/CreateQuestion'

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
            setList([])
            let res = await server.ListQuestions(search)
            setList(res.data)
        } catch (err) {
            console.log("err: ", err)
        }
    }

    async function create(data) {
        try {
            let res = await server.CreateQuestion(data)
            console.log("createQuestions >>> ", res)
            setList([])
            searchListQuestions("")
        } catch (err) {
            console.log("err: ", JSON.stringify(err))
        }
    }

    return (
        <>
            <div className="question">
                <CreateQuestion onSubmit={create} />
            </div>
            <div>
                <Search onSubmit={searchListQuestions} />
                <ListQuestion data={list} />
            </div>

        </>
    );
}
