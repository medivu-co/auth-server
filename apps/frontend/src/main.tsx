import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./index.css";
import AuthorizePage from "./pages/Authorize";
import NotFoundPage from "./pages/NotFound";
import SignUpPage from "./pages/SignUp";

// biome-ignore lint/style/noNonNullAssertion: <alppano>
createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/authorize" element={<AuthorizePage />} />
        {/* path="/about"은 소개 페이지 경로 */}
        <Route path="/signup" element={<SignUpPage />} />
        {/* 정의되지 않은 다른 모든 경로 처리 (예: 404 페이지) */}
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>
);
