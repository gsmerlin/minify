import { Button } from "@mui/material";
import { googleLogout } from "@react-oauth/google";
import React from "react";
import { Page, usePageActions } from "../../atoms/page";
import { BLANK_USER, useUserActions } from "../../atoms/user";

export const Logout: React.FC = () => {
  const { setUserInfo } = useUserActions();
  const { setPage } = usePageActions();
  const handleClick = () => {
    setUserInfo(BLANK_USER);
    setPage(Page.Login);
    googleLogout();
  };

  return <Button onClick={() => handleClick()}>Logout</Button>;
};
