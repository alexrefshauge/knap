import { useQuery } from "@tanstack/react-query"
import axios from "axios"
import formatDate from "../utils/dateFormat"
import WeekDayStats from "./WeekDayStats"

import("./StatsPage.css")

type Response = string[][]

export default function StatsPage() {
    const todayDate = new Date()
    const params = new URLSearchParams([["date", formatDate(todayDate)]])
    const { data: week, isLoading, isError } = useQuery<Response>({
        queryKey: ["count", "week"],
        queryFn: async () => (await axios.get("/press/week", {
            params: params
        })).data,
    })

    if (isLoading) return <p>loading...</p>
    if (isError) return <p>error</p>

    return <div className="stats-page">{week?.map(presses => {
        return <WeekDayStats presses={presses} />
    })}</div>
}