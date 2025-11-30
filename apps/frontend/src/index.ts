import { serve } from 'bun';
import index from './index.html';

const server = serve({
  development: process.env.NODE_ENV !== 'production' && {
    // Echo console logs from the browser to the server
    console: true,
    // Enable browser hot reloading in development
    hmr: true,
  },
  port: 5173,
  routes: {
    // Serve index.html for all unmatched routes.
    '/*': index,
  },
});

console.log(`🚀 Server running at ${server.url}`);
