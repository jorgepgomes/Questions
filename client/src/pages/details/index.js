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

    async function AnswerQuestion() {

    }

    return (
        <>
            <div className="details">
                <button type="button" onClick={goBack} className="buttonBack">voltar</button>
                <div className="question">
                    <strong>{question.Text}</strong>
                </div>
                <div>
                    <Answers data={question.Answers} />
                </div>
                <div>
                    <AnswerForm onSubmit={AnswerQuestion} />
                </div>
            </div>
        </>
    );
}
