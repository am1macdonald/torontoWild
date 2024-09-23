const login = async () => {
    const app = document.getElementById('app')
    if (!app) {
        throw new Error('App element not found')
    }
    app.innerHTML = `
        <h1>Login</h1>
        <form>
            <label for="username">Username</label>
            <input type="text" name="username" id="username">
            <label for="password">Password</label>
            <input type="password" name="password" id="password">
            <button type="submit">Login</button>
        </form>
    `
}

export default login
