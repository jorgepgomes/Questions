import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom'
import server from "../../services/server"
import Answers from '../../components/Answers';
import AnswerForm from '../../components/AnswerForm'
import './style.css'


export default function Details() {
    const { pathname } = useLocation()

    const [question, setQuestion] = useState({});

    useEffect(() => {
        function getId() {
            let id = pathname.replace("/", "")
            return id
        }

        async function getData(id) {
            try {
                let result = await server.DetailsQuestion(id)
                setQuestion(result.data)
            } catch (err) {
                console.log(err)
            }
        }

        let id = getId()
        getData(id)

    }, [pathname])

    function goBack() {
        window.location.href = `/`;
    }

    function getId() {
        let id = pathname.replace("/", "")
        return id
    }

    async function getData(id) {
        try {
            let result = await server.DetailsQuestion(id)
            setQuestion(result.data)
        } catch (err) {
            console.log(err)
        }
    }

    async function AnswerQuestion(data) {
        try {
            let id = getId()
            let res = await server.AnswerQuestion(id, data)
            console.log(res)
            getData(id)
        } catch (err) {
            console.log("err ", err)
        }
    }

    async function likeOrDislike(answerId, type, value) {
        try {
            let id = getId()
            let body = {
                id_question: parseInt(id),
                id_answer: answerId,
                local: type,
                like: value
            }
            let res = await server.Like(body)
            console.log(res)
            getData(id)
        } catch (err) {
            console.log(err)
        }
    }

    return (
        <>
            <div className="details">
                <button type="button" onClick={goBack} className="buttonBack">voltar</button>
                <div className="question">
                    <strong>{question.Text}</strong>
                    <div className="react-answer">
                        <strong className="action" style={{marginRight: "10px"}}>likes {question.Likes}</strong>
                        <strong className="action" onClick={() => likeOrDislike(0, "question", 1)}>Like</strong>
                        <strong className="action" onClick={() => likeOrDislike(0, "question", -1)}>Dislike</strong>
                    </div>
                </div>
                <div>
                    <Answers data={question.Answers} reactionAnswer={likeOrDislike} />
                </div>
                <div>
                    <AnswerForm onSubmit={AnswerQuestion} />
                </div>
            </div>
        </>
    );
}
