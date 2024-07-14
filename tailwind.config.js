const plugin = require("tailwindcss/plugin");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/**/*.{html,js}", "./files/assets/**/*.{css}"],
  theme: {
    extend: {},
  },
  plugins: [
    require('daisyui'),
  ],
};