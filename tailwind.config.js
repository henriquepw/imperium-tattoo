
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

function generateAlphaScale(name) {
  return {
    1: `var(--${name}-a1)`,
    2: `var(--${name}-a2)`,
    3: `var(--${name}-a3)`,
    4: `var(--${name}-a4)`,
    5: `var(--${name}-a5)`,
    6: `var(--${name}-a6)`,
    7: `var(--${name}-a7)`,
    8: `var(--${name}-a8)`,
    9: `var(--${name}-a9)`,
    10: `var(--${name}-a10)`,
    11: `var(--${name}-a11)`,
    12: `var(--${name}-a12)`,
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
        agray: generateAlphaScale("sand"),
        info: generateScale("cyan"),
        error: generateScale("tomato"),
        success: generateScale("green"),
        warning: generateScale("amber"),
      },
    },
    fontFamily: {
      sans: ["raleway"],
    },
  },
  plugins: [],
}

