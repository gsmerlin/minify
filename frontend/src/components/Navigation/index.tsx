import React from "react";
import { Page, usePage } from "../../atoms/page";
import { Delete } from "../Delete";
import { Home } from "../Home";
import { Login } from "../Login";

export const Navigation: React.FC = () => {
  const page = usePage();
  switch (page) {
    case Page.Home:
      return <Home />;
    case Page.Login:
      return <Login />;
    case Page.Delete:
      return (
        <>
          <Delete />
          <Home />
        </>
      );
    default:
      return <div>Unknown</div>;
  }
};
