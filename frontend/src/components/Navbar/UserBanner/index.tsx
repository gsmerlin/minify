import { Box, Avatar } from "@mui/material";
import { useGetUserInfo } from "../../../atoms/user";
import { Logout } from "../../Logout";

export const UserBanner: React.FC = () => {
  const { picture } = useGetUserInfo();
  if (picture === "") {
    return null;
  }
  return (
    <Box
      width="100%"
      height="100%"
      display="flex"
      alignItems="center"
      justifyContent="flex-end"
    >
      <Box display="inherit" borderRadius="5px" padding="10px">
        <Box display="flex" alignItems="center">
          <Avatar src={picture} />
        </Box>
        <Logout />
      </Box>
    </Box>
  );
};
