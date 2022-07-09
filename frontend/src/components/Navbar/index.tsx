import { Box } from "@mui/material";
import React from "react";
import { Logo } from "./Logo";
import { UserBanner } from "./UserBanner";

export const Navbar: React.FC = () => {
  return (
    <Box
      width="100%"
      height="5%"
      display="inline-flex"
      alignItems="center"
      marginBottom="1rem"
      borderBottom="1px solid black"
      paddingBottom="10px"
    >
      <Logo />
      <UserBanner />
    </Box>
  );
};
