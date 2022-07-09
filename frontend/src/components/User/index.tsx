import React from "react";
import { useGetUserInfo } from "../../atoms/user";
import { Logout } from "../Logout";
import { Info } from "./styled";

export const User: React.FC = () => {
  const { name, picture } = useGetUserInfo();
  return (
    <Info>
      <img src={picture} alt={name} />
      {name}
      <Logout />
    </Info>
  );
};
