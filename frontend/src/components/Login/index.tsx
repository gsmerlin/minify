import React from "react";
import jwt_decode from "jwt-decode";
import { useUserActions } from "../../atoms/user";
import { CredentialResponse, GoogleLogin } from "@react-oauth/google";
import { Box, Typography } from "@mui/material";
import { Page, usePageActions } from "../../atoms/page";

type GoogleJWT = {
  sub: string;
  iat: number;
  exp: number;
  given_name: string;
  family_name: string;
  name: string;
  email: string;
  picture: string;
};

export const Login: React.FC = () => {
  const { setUserInfo } = useUserActions();
  const { setPage } = usePageActions();
  const onSuccess = (response: CredentialResponse) => {
    const userObject = jwt_decode<GoogleJWT>(response.credential!);
    setUserInfo({
      token: userObject.sub,
      name: userObject.given_name,
      email: userObject.email,
      picture: userObject.picture,
    });
    setPage(Page.Home);
  };

  const onError = () => {
    console.log("Login failed");
  };

  return (
    <Box
      display="flex"
      flexDirection="column"
      alignItems="center"
      marginTop="100px"
      minHeight="100vh"
    >
      <Box
        border="1px solid black"
        height="250px"
        width="600px"
        display="inherit"
        flexDirection="column"
        alignItems="center"
        justifyContent="center"
        borderRadius="50px"
      >
        <Typography component="div" variant="h2">
          Welcome to Minify!
        </Typography>
        <Typography component="span" variant="h6" marginBottom="5px">
          Please log in using google to continue:
        </Typography>
        <GoogleLogin onSuccess={onSuccess} onError={onError} />
      </Box>
    </Box>
  );
};
