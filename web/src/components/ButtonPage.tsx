import { useRef, useState } from 'react';
import './ButtonPage.css';
import { useMutation } from '@tanstack/react-query';
import axios from 'axios';

function ButtonPage() {
    const [text, setText] = useState("PRESS ME")
    const pressMutation = useMutation({
        mutationFn: () => axios.put('/press'),
        onSuccess: () => {
            buttonRef.current?.classList.remove('pressed');
        },
        onError: (err) => {
            buttonRef.current?.classList.remove('pressed')
            setText(err.message)
        }
    })
    const buttonRef = useRef<HTMLDivElement>(null);

    const handleClick = () => {
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