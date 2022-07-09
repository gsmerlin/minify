import { Box, Typography } from "@mui/material";
import { useLinks } from "../atoms";
import { LinkItem } from "./LinkItem";

export const LinkList = () => {
  const data = useLinks();

  return (
    <Box marginTop="1rem" width="100%">
      {data.length > 0 &&
        data.map((link) => <LinkItem key={link.ID} link={link} />)}
      {data.length === 0 && (
        <Box display="flex" alignItems="center" flexDirection="column">
          <Typography component="p" variant="h5">
            You have no links. Create one using the buttom below!
          </Typography>
        </Box>
      )}
    </Box>
  );
};
