export default function formatDate(d: Date): string {
    return `${d.getDate()}-${d.getMonth()+1}-${d.getFullYear()}`
}