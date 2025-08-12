import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
import sveltePreprocess from 'svelte-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://svelte.dev/docs#compile-time-svelte-preprocess
	// for more information about preprocessors
	preprocess: [
		vitePreprocess(),
		sveltePreprocess({
			sass: {
				// Use indented syntax and prepend global styles
				indentedSyntax: true,
				prependData: `@import 'src/lib/styles/variables.sass'`
			},
			pug: true
		})
	],

	kit: {
		adapter: adapter()
	}
};

export default config;
