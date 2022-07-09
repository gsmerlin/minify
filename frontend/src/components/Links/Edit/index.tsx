import {
  Button,
  DialogActions,
  DialogContent,
  DialogContentText,
  FormHelperText,
  TextField,
} from "@mui/material";
import React from "react";
import { useState } from "react";
import { useCallPopup } from "../../../atoms/popup";
import { Link, useLinkActions, useSelectedLink } from "../atoms";
import { validateURL } from "../Create/helpers";

export const Edit = () => {
  const selectedLink = useSelectedLink();
  const { editLink, fetchLinks } = useLinkActions();
  const callPopup = useCallPopup();
  const [link, setLink] = useState<Link>(
    selectedLink || {
      ID: "",
      Destination: "",
      Email: "",
    }
  );
  const [error, setError] = React.useState("");

  const closePopup = () => {
    callPopup({ close: true });
  };

  const handleEdit = async () => {
    if (link.Destination === selectedLink?.Destination) {
      closePopup();
      return;
    }
    const errorMsg = await validateURL(link.Destination);
    if (errorMsg) {
      setError(errorMsg);
      return;
    }
    editLink(link);
    fetchLinks();
    closePopup();
  };
  return (
    <DialogContent>
      <DialogContentText>Edit link {link.ID}</DialogContentText>
      <TextField
        margin="dense"
        id="destination"
        label="Destination"
        type="text"
        fullWidth
        value={link.Destination}
        onChange={(e) => setLink({ ...link, Destination: e.target.value })}
      />
      <FormHelperText>{error}</FormHelperText>
      <DialogActions>
        <Button onClick={() => closePopup()}>Cancel</Button>
        <Button onClick={() => handleEdit()}>Save</Button>
      </DialogActions>
    </DialogContent>
  );
};
