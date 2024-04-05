const colors = require("tailwindcss/colors");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["cmd/web/*.templ"],
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};
