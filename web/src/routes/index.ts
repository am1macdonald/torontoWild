const home = async () => {
    const app = document.getElementById('app')
    if (!app) {
        throw new Error('App element not found')
    }
    app.innerHTML = `
        <h1>Home</h1>
        <a href="/login">login</a>
    `
}

export default home
