/**
 * This file is the entry point for the React app, it sets up the root
 * element and renders the App component to the DOM.
 *
 * It is included in `src/index.html`.
 */

import { createRoot } from 'react-dom/client';
import { BrowserRouter, Route, Routes } from 'react-router';
import { ThemeProvider } from './components/theme-provider';
import { AuthorizePage } from './pages/authorize';
import { ErrorPage } from './pages/error';

import './styles/globals.css';
import { SignUpPage } from './pages/signup';

// biome-ignore lint/style/noNonNullAssertion: <bun default code>
const elem = document.getElementById('root')!;
const app = (
  <ThemeProvider>
    <BrowserRouter>
      <Routes>
        <Route element={<AuthorizePage />} path='/authorize' />
        <Route element={<SignUpPage />} path='/signup' />
        <Route element={<ErrorPage />} path='*' />
      </Routes>
    </BrowserRouter>
  </ThemeProvider>
);

if (import.meta.hot) {
  // With hot module reloading, `import.meta.hot.data` is persisted.
  // biome-ignore lint/suspicious/noAssignInExpressions: <bun default code>
  const root = (import.meta.hot.data.root ??= createRoot(elem));
  root.render(app);
} else {
  // The hot module reloading API is not available in production.
  createRoot(elem).render(app);
}
