import { Typography, Box } from "@mui/material";
import React, { Suspense } from "react";
import { useGetUserInfo } from "../../atoms/user";
import { LinkList } from "./LinkList";

export const Home: React.FC = () => {
  const { name } = useGetUserInfo();
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
      </Box>
      <LinkList />
    </Suspense>
  );
};
