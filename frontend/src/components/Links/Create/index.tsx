import {
  Button,
  DialogActions,
  DialogContent,
  FormHelperText,
  TextField,
  Typography,
} from "@mui/material";
import React from "react";
import { useCallPopup } from "../../../atoms/popup";
import { useLinkActions } from "../atoms";
import { validateURL } from "./helpers";

export const Create: React.FC = () => {
  const callPopup = useCallPopup();
  const { createLink } = useLinkActions();
  const [linkUrl, setLinkUrl] = React.useState("");
  const [error, setError] = React.useState("");

  const closePopup = () => {
    callPopup({ close: true });
  };

  const handleCreate = async () => {
    const errorMsg = await validateURL(linkUrl);
    if (errorMsg) {
      setError(errorMsg);
      return;
    }
    createLink(linkUrl);
    closePopup();
  };

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setLinkUrl(event.target.value);
  };
  return (
    <DialogContent>
      <Typography
        component="h1"
        variant="caption"
        sx={{ marginBottom: "10px" }}
      >
        Create new Link
      </Typography>
      <TextField
        margin="dense"
        value={linkUrl}
        onChange={handleChange}
        variant="outlined"
        label="https://... "
      ></TextField>
      <FormHelperText>{error}</FormHelperText>
      <DialogActions>
        <Button onClick={() => closePopup()}>Cancel</Button>
        <Button onClick={async () => await handleCreate()}>Create</Button>
      </DialogActions>
    </DialogContent>
  );
};
