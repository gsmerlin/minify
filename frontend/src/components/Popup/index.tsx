import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
} from "@mui/material";
import React from "react";
import { usePopupValues } from "../../atoms/popup";

export const Popup: React.FC = () => {
  const { title, content, open, actions } = usePopupValues();
  console.log("I was called!", open);
  return (
    <Dialog open={open}>
      <DialogTitle>{title}</DialogTitle>
      <DialogContent>
        <DialogContentText>{content}</DialogContentText>
      </DialogContent>
      <DialogActions>
        {actions.map((action) => (
          <Button key={action.name} onClick={action.callback}>
            {action.name}
          </Button>
        ))}
      </DialogActions>
    </Dialog>
  );
};
