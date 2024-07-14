const plugin = require("tailwindcss/plugin");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/**/*.{html,js}", "./files/assets/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [
    require("flowbite-typography"),
  ],
};