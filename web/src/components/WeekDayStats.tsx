import { useEffect, useRef } from "react"



export default function ({ presses }: { presses: string[] }) {
    const canvasRef = useRef<HTMLCanvasElement>(null)
    const pressTimes = presses.map(p => new Date(p))

    useEffect(() => {
        const canvas = canvasRef.current
        if (!canvas) return
        const ctx = canvas.getContext('2d')
        if (!ctx) return
        //Our first draw
        ctx.fillStyle = 'rgb(219, 209, 204)'
        ctx.fillStyle = 'rgb(247, 240, 224)'
        ctx.strokeStyle = 'rgb(236, 65, 42)'

        const height = ctx.canvas.height
        pressTimes.forEach(t => {
            const xFrac = ((t.getHours() * 60 * 60) + (t.getMinutes() * 60) + t.getSeconds()) / (24 * 60 * 60)
            const x = xFrac * ctx.canvas.width

            ctx.beginPath()
            ctx.moveTo(x, 0)
            ctx.lineTo(x, height)
            ctx.stroke()
        })
    }, [])

    return <div className="weekday-stats">
        <h2>{presses.length}</h2>
        <canvas ref={canvasRef}></canvas>
    </div>
}