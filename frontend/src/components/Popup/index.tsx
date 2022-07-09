import { Dialog } from "@mui/material";
import React from "react";
import { usePopupValues } from "../../atoms/popup";

export const Popup: React.FC = () => {
  const { content, open } = usePopupValues();
  return <Dialog open={open}>{content}</Dialog>;
};
