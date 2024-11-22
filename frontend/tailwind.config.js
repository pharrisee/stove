/** @type {import('tailwindcss').Config} */
import forms from "@tailwindcss/forms";
import typography from "@tailwindcss/typography";

const purple = {
  50: "#8E4490",
  100: "#6F3571",
  200: "#5A2B5C",
  300: "#512753",
  400: "#452146",
  500: "#391F3A",
  600: "#3A193F",
  700: "#34163A",
  800: "#2A122E",
  900: "#1E0D21",
};

const navy = {
  50: "#374A6E",
  100: "#324467",
  200: "#2F4060",
  300: "#2B3A58",
  400: "#28354E",
  500: "#212B40",
  600: "#172031",
  700: "#111724",
  800: "#0B1019",
  900: "#090C12",
};

export default {
  content: [
    "../web/**/*.templ",
    "../web/**/*.go",
  ],
  theme: {
    extend: {
      colors: {
        black: "#383735",
        purple: purple,
        navy: navy,
        cu_navy: {
          200: "#394563",
          300: "#3D4C6D",
          400: "#2b354c",
          500: "#242C3F",
        },
        cu_red: {
          400: "#d4374a",
        },
        cu_link: {
          400: "#045bc6",
          500: "#022d62",
        },
        cu_black: {
          400: "#383735",
        },
        cu_blue: {
          400: "#3a6cb5",
        },
        // CU brand colors, 400 base colors, all other are shades
        "brand-black": "#121212",
        "brand-white": "#FFFFFF",
        "brand-red": {
          400: "#E4251B",
          410: "#B61E16",
        },
        "brand-grey": {
          800: "#333333",
          700: "#4C4C4C",
          600: "#666666",
          500: "#808080",
          400: "#999999",
          300: "#B3B3B3",
          200: "#CCCCCC",
          100: "#E5E5E5",
          50: "#F2F2F2",
        },
        "brand-link-blue": {
          500: "#034494",
          400: "#045BC6",
          300: "#0572F8",
        },
        "brand-orange": {
          500: "#C15E13",
          400: "#E9761E",
          300: "#EE924C",
        },
        "brand-yellow": {
          500: "#CC8F00",
          400: "#FFB300",
          300: "#FEF6DA",
        },
        "brand-cadet": {
          500: "#459F81",
          400: "#5EB99B",
          300: "#82C8B1",
        },
        "brand-yellow-green": {
          500: "#6B9B33",
          400: "#85C041",
          300: "#9ECD67",
        },
        "brand-forest-green": {
          500: "#045628",
          400: "#07873E",
          300: "#0AB754",
        },
        "brand-light-blue": {
          500: "#0077EA",
          400: "#1E90FF",
          300: "#51A9FF",
        },
        "brand-royal-blue": {
          500: "#190FFF",
          400: "#4A42FF",
          300: "#7B75FF",
        },
        "brand-midnight-blue": {
          500: "#0A1A3D",
          400: "#273573",
          300: "#344799",
        },
        "brand-indigo": {
          500: "#360152",
          400: "#570185",
          300: "#7801B8",
        },
        "brand-violet": {
          500: "#A10196",
          400: "#D401C5",
          300: "#FE0AED",
        },
        "brand-stone": {
          500: "#B3A99D",
          400: "#C9C2BA",
          300: "#DFDBD7",
        },
      },
      fontFamily: {
        bold: [
          "franklin_gothic_fs_medregular",
          "franklin_gothic_fs_bookRg",
          '"Helvetica Neue"',
          "system-ui",
          "Arial",
        ],
        FranklinGothicURWLig: ["FranklinGothicURW-Lig"],
        FranklinGothicBookRg: ["franklin_gothic_fs_bookRg"],
        FranklinLight: ["tee_franklin_light"],
        FranklinGothicMedregular: ["franklin_gothic_fs_medregular"],
      }
    },
  },
  plugins: [forms, typography, require('tailwindcss-animated')],
};
