import {
  Button,
  DialogActions,
  DialogContent,
  DialogContentText,
} from "@mui/material";
import { useCallPopup } from "../../../atoms/popup";
import { useLinkAnalytics } from "../atoms";

export const Analytics = () => {
  const analytics = useLinkAnalytics();
  const callPopup = useCallPopup();
  const closePopup = () => callPopup({ close: true });
  return (
    <DialogContent>
      <DialogContentText>
        Analytics for www.minify.com/{analytics.ID}
      </DialogContentText>
      <DialogContentText>
        Destination: {analytics.Destination}
      </DialogContentText>
      <DialogContentText>
        Total Clicks: {analytics.TotalClicks}
      </DialogContentText>
      <DialogContentText>Timestamps of clicks: </DialogContentText>
      {analytics.Timestamps.map((timestamp, index) => (
        <DialogContentText key={index}>{timestamp}</DialogContentText>
      ))}
      <DialogActions>
        <Button onClick={() => closePopup()}>Done</Button>
      </DialogActions>
    </DialogContent>
  );
};
