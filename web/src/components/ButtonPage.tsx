import { useRef, useState } from 'react';
import './ButtonPage.css';
import { useMutation } from '@tanstack/react-query';
import axios from 'axios';

function ButtonPage() {
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
        },
        onError: (err) => {
            resetButton()
            setText(err.message)
        }
    })

    const handleClick = () => {
        if (pressed) return
        setPressed(true)
        buttonRef.current?.classList.add('pressed');
        pressMutation.mutate()
    }

    return (
        <div className='button-page' >
            <div className='button' >
                <div className='button-top' onClick={handleClick} >
                    <p className='button-text'>{pressMutation.isPending ? "loading..." : text}</p>
                </div>
                <div className='button-middle' onClick={handleClick} ref={buttonRef} />
                <div className='button-bottom' />
            </div>
        </div >
    );
}

export default ButtonPage;