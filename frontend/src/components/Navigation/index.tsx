import React from "react";
import { usePage } from "../../atoms/page";
import { Home } from "../Home";
import { Login } from "../Login";

export const Navigation: React.FC = () => {
  const page = usePage();

  switch (page) {
    case "home":
      return <Home />;
    case "login":
      return <Login />;
    default:
      return <div>Unknown</div>;
  }
};
