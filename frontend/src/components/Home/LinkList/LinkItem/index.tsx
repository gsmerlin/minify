import { Box, Button, Tooltip, Typography } from "@mui/material";
import { Edit, Delete, Visibility } from "@mui/icons-material";
import { Page, usePageActions } from "../../../../atoms/page";
import { Link, useLinkActions } from "../../atoms";

interface ILinkProps {
  link: Link;
}
export const LinkItem = ({ link }: ILinkProps) => {
  const { setPage } = usePageActions();
  const { setSelectedLink } = useLinkActions();

  const handlePageChange = (page: Page, id: string) => {
    setSelectedLink(id);
    setPage(page);
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
            <Button onClick={() => handlePageChange(Page.Analytics, link.ID)}>
              <Visibility />
            </Button>
          </Tooltip>
          <Tooltip title="Edit Link">
            <Button onClick={() => handlePageChange(Page.Edit, link.ID)}>
              <Edit />
            </Button>
          </Tooltip>
          <Tooltip title="Delete Link">
            <Button onClick={() => handlePageChange(Page.Edit, link.ID)}>
              <Delete />
            </Button>
          </Tooltip>
        </Box>
      </Box>
    </Box>
  );
};
