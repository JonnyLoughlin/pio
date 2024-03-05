// prettier.config.js, .prettierrc.js, prettier.config.mjs, or .prettierrc.mjs

/** @type {import("prettier").Config} */
const config = {
	printWidth: 120,
	tabWidth: 4,
	singleAttributePerLine: true,
	plugins: ["prettier-plugin-tailwindcss", "prettier-plugin-go-template"],
	goTemplateBracketSpacing: true,
};

export default config;
