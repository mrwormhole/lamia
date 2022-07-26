
type BigMessage = {
    message: string
}

export default function Alert() {
    let m : BigMessage = {
        message: "Hello world"
    }
    alert(m.message)
}