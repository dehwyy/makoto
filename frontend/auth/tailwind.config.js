import defaultTheme from 'tailwindcss/defaultTheme';

/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontFamily: {
				Jua: ['Jua', ...defaultTheme.fontFamily.sans],
				Content: ['Comfortaa Variable', ...defaultTheme.fontFamily.sans],
				Kanji: ['Klee One', ...defaultTheme.fontFamily.sans],
				ContentT: ['Shantell Sans Variable', ...defaultTheme.fontFamily.sans]
			}
		}
	},
	daisyui: {
		themes: [
			{
				mytheme: {
					primary: '#ff00d2',

					secondary: '#00a2ff',

					accent: '#4629f2',

					neutral: '#2a323c',

					'base-100': '#1d232a',

					info: '#7583ca',

					success: '#36d399',

					warning: '#fbbd23',

					error: '#f87272'
				}
			}
		]
	},
	plugins: [require('daisyui')]
};
