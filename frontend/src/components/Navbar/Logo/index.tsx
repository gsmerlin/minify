import React from "react";
import { Typography } from "@mui/material";
import DiamondIcon from "@mui/icons-material/Diamond";

export const Logo: React.FC = () => (
  <Typography variant="h5" component="div" height="100%" width="10%">
    <DiamondIcon
      sx={{
        width: "30px",
        height: "30px",
        verticalAlign: "middle",
      }}
    />
    Minify
  </Typography>
);
