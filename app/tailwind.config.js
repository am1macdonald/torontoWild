/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",
    ],
    plugins: [
        require("@tailwindcss/typography"),
        require('daisyui'),
    ],
    daisyui: {
        themes: ['nord', 'sunset'],
    },
    darkMode: ['selector', '[data-theme="sunset"]']
}

