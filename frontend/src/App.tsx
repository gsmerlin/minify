import { GoogleOAuthProvider } from "@react-oauth/google";
import React, { Suspense } from "react";
import { Navigation } from "./components/Navigation";
import { QueryClient, QueryClientProvider } from "react-query";
import { Navbar } from "./components/Navbar";

const clientId =
  "1072806869736-dirrf14sfk8u03o6a0jchh47hglaljs1.apps.googleusercontent.com";

const queryClient = new QueryClient();

const App: React.FC = () => (
  <Suspense fallback="Loading...">
    <GoogleOAuthProvider clientId={clientId}>
      <QueryClientProvider client={queryClient}>
        <Navbar />
        <Navigation />
      </QueryClientProvider>
    </GoogleOAuthProvider>
  </Suspense>
);
export default App;
