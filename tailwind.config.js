
function generateScale(name) {
  return {
    1: `var(--${name}-1)`,
    2: `var(--${name}-2)`,
    3: `var(--${name}-3)`,
    4: `var(--${name}-4)`,
    5: `var(--${name}-5)`,
    6: `var(--${name}-6)`,
    7: `var(--${name}-7)`,
    8: `var(--${name}-8)`,
    9: `var(--${name}-9)`,
    10: `var(--${name}-10)`,
    11: `var(--${name}-11)`,
    12: `var(--${name}-12)`,
  };
}

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./web/view/**/*.templ",
  ],
  theme: {
    extend: {
      colors: {
        accent: generateScale("orange"),
        gray: generateScale("sand"),
      }
    },
    fontFamily: {
      sans: ["raleway"],
    },
  },
  plugins: [],
}

