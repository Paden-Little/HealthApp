import type { Config } from 'tailwindcss';
import flowbite from 'flowbite/plugin';

export default {
  plugins: [flowbite],
  content: ['./node_modules/flowbite/**/*.{js,ts}'],
};
