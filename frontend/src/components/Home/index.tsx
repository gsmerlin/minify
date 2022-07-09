import { Typography, Box, Button } from "@mui/material";
import React, { Suspense } from "react";
import { useCallPopup } from "../../atoms/popup";
import { useGetUserInfo } from "../../atoms/user";
import { Create } from "../Links/Create";
import { LinkList } from "../Links/LinkList";

export const Home: React.FC = () => {
  const { name } = useGetUserInfo();
  const callPopup = useCallPopup();
  return (
    <Suspense fallback="Loading...">
      <Box
        display="flex"
        flexDirection="column"
        alignItems="center"
        marginTop="3%"
      >
        <Typography component="span" variant="h5">
          Welcome to Minify, {name}!
        </Typography>
        <LinkList />
        <Button onClick={() => callPopup({ content: <Create /> })}>
          + Create New Link
        </Button>
      </Box>
    </Suspense>
  );
};
