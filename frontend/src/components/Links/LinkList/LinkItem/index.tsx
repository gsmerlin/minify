import { Box, Button, Tooltip, Typography } from "@mui/material";
import {
  Edit as EditIcon,
  Delete as DeleteIcon,
  Visibility as AnalyticsIcon,
} from "@mui/icons-material";
import { Link, useLinkActions } from "../../atoms";
import { useCallPopup } from "../../../../atoms/popup";
import { Delete } from "../../Delete";
import React from "react";
import { Analytics } from "../../Analytics";
import { Edit } from "../../Edit";

interface ILinkProps {
  link: Link;
}
export const LinkItem = ({ link }: ILinkProps) => {
  const { setSelectedLink } = useLinkActions();
  const callPopup = useCallPopup();

  const handleButtons = (id: string, content: React.ReactNode) => {
    setSelectedLink(id);
    callPopup({ content });
  };

  return (
    <Box display="flex" width="100%" marginBottom="5px" key={link.ID}>
      <Box
        border="1px solid black"
        width="99%"
        padding="10px"
        display="inline-flex"
      >
        <Box
          width="85%"
          marginLeft="15%"
          display="flex"
          flexDirection="column"
          alignItems="center"
          justifyContent="flex-end"
        >
          <Typography component="p" fontWeight="bold" variant="h5">
            www.minify.com/{link.ID}
          </Typography>
          <Typography component="span" variant="subtitle1">
            {link.Destination}
          </Typography>
        </Box>
        <Box width="15%" display="inline-flex">
          <Tooltip title="View Analytics">
            <Button onClick={() => handleButtons(link.ID, <Analytics />)}>
              <AnalyticsIcon />
            </Button>
          </Tooltip>
          <Tooltip title="Edit Link">
            <Button onClick={() => handleButtons(link.ID, <Edit />)}>
              <EditIcon />
            </Button>
          </Tooltip>
          <Tooltip title="Delete Link">
            <Button onClick={() => handleButtons(link.ID, <Delete />)}>
              <DeleteIcon />
            </Button>
          </Tooltip>
        </Box>
      </Box>
    </Box>
  );
};
