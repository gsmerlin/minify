import React from "react";
import { Page, usePage } from "../../atoms/page";
import { Home } from "../Home";
import { Login } from "../Login";

export const Navigation: React.FC = () => {
  const page = usePage();
  switch (page) {
    case Page.Home:
      return <Home />;
    case Page.Login:
      return <Login />;
    default:
      return <div>Unknown</div>;
  }
};
