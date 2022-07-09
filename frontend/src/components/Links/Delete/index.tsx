import {
  DialogContent,
  Typography,
  DialogActions,
  Button,
} from "@mui/material";
import { useCallPopup } from "../../../atoms/popup";
import { useLinkActions, useSelectedLink } from "../atoms";

export const Delete: React.FC = () => {
  const { deleteLink } = useLinkActions();
  const selectedLink = useSelectedLink();
  const callPopup = useCallPopup();

  const closePopup = () => {
    callPopup({ close: true });
  };

  const handleDelete = async () => {
    if (selectedLink) {
      await deleteLink(selectedLink.ID);
      closePopup();
    }
  };

  return (
    <DialogContent>
      <Typography
        component="h1"
        variant="caption"
        sx={{ marginBottom: "10px" }}
      >
        Delete Link
      </Typography>
      <Typography component="p" variant="h5">
        Are you sure you want to delete this link?
      </Typography>
      <DialogActions>
        <Button onClick={() => closePopup()}>Cancel</Button>
        <Button onClick={async () => await handleDelete()}>Delete</Button>
      </DialogActions>
    </DialogContent>
  );
};
