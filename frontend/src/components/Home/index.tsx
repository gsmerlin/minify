import { Typography, Box } from "@mui/material";
import React, { Suspense } from "react";
import { useGetUserInfo } from "../../atoms/user";
import { useLinks } from "./atoms";

export const Home: React.FC = () => {
  const { name } = useGetUserInfo();
  const data = useLinks();
  return (
    <Suspense fallback="...loading">
      <Box display="flex" flexDirection="column" alignItems="center">
        <Typography component="span" variant="h5">
          Welcome to Minify, {name}!
        </Typography>
      </Box>
      <Box marginTop="1rem" width="100%">
        {data.length > 0 &&
          data.map((link) => (
            <Box display="flex" width="100%" marginBottom="5px" key={link.ID}>
              <Box border="1px solid black" width="99%" padding="10px">
                <Typography component="p" fontWeight="bold" variant="h5">
                  www.minify.com/{link.ID}
                </Typography>
                <Typography component="span" variant="subtitle1">
                  {link.Destination}
                </Typography>
              </Box>
            </Box>
          ))}
        {data.length === 0 && (
          <Typography component="p" variant="h5">
            You have no links. Create one using the buttom below!
          </Typography>
        )}
      </Box>
    </Suspense>
  );
};
