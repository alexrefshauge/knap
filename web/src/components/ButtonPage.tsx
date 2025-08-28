import { useRef, useState } from 'react';
import './ButtonPage.css';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import axios from 'axios';
import formatDate from '../utils/dateFormat';

export default function ButtonPage() {
    const client = useQueryClient()
    const [pressed, setPressed] = useState(false)
    const buttonRef = useRef<HTMLDivElement>(null);
    const resetButton = () => setTimeout(() => {
        setPressed(false)
        setText(":(")
        buttonRef.current?.classList.remove('pressed')
    }, 500)
    const [text, setText] = useState(":)")
    const pressMutation = useMutation({
        mutationFn: () => axios.put('/press'),
        onSuccess: () => {
            resetButton()
            client.invalidateQueries({ queryKey: ["count", "today"] })
        },
        onError: (err) => {
            resetButton()
            setText(err.message)
        }
    })

    const todayDate = new Date()
    const countParams = new URLSearchParams([["count", "t"], ["date", formatDate(todayDate)]])
    const { data: today, isLoading, isError } = useQuery<{ count: number }>({
        queryKey: ["count", "today"],
        queryFn: async () => (await axios.get("/press/today", {
            params: countParams
        })).data,
    })

    const handleClick = () => {
        if (pressed) return
        setPressed(true)
        buttonRef.current?.classList.add('pressed');
        pressMutation.mutate()
    }

    if (isLoading) return <p>loading</p>
    if (isError) return <p>error</p>

    return (
        <div className='button-page' >
            <div className='button' >
                <div className='button-top' onClick={handleClick} >
                    <p className='button-text'>{pressMutation.isPending ? "loading..." : text}</p>
                </div>
                <div className='button-middle' onClick={handleClick} ref={buttonRef} />
                <div className='button-bottom' />
            </div>
            <div className='press-count'><h1>Count Today: {today?.count}</h1></div>
        </div >
    );
}

