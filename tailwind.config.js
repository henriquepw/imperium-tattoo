/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./static/**/*",
    "!./static/css/output.",
    "./view/**/*.{html,js,templ}",
  ],
  theme: {
    extend: {},
    fontFamily: {
      sans: ["raleway"],
    },
  },
  plugins: [],
}

