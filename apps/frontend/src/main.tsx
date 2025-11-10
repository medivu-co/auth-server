import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import "./index.css";
import AuthorizePage from "./pages/Authorize";
import ErrorPage from "./pages/Error";
import SignUpPage from "./pages/SignUp";

// biome-ignore lint/style/noNonNullAssertion: <alppano>
createRoot(document.getElementById("root")!).render(
	<StrictMode>
		<BrowserRouter>
			<Routes>
				<Route
					path="/"
					element={
						<Navigate
							to="/authorize?response_type=code&client_id=test&redirect_uri=http%3A%2F%2F127.0.0.1%3A3000%2Fredirect"
							replace
						/>
					}
				/>
				<Route path="/authorize" element={<AuthorizePage />} />
				<Route path="/signup" element={<SignUpPage />} />
				{/* 404 Route */}
				<Route path="*" element={<ErrorPage />} />
			</Routes>
		</BrowserRouter>
	</StrictMode>,
);
