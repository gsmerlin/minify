import { Box, Typography } from "@mui/material";
import { useFetchLinks } from "../atoms";
import { LinkItem } from "./LinkItem";

export const LinkList = () => {
  const data = useFetchLinks();

  return (
    <Box marginTop="1rem" width="100%">
      {data.length > 0 &&
        data.map((link) => <LinkItem key={link.ID} link={link} />)}
      {data.length === 0 && (
        <Typography component="p" variant="h5">
          You have no links. Create one using the buttom below!
        </Typography>
      )}
    </Box>
  );
};
