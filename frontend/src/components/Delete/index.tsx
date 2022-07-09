import { useCallPopup } from "../../atoms/popup";
import { useSelectedLink } from "../Home/atoms";

export const Delete = () => {
  const callPopup = useCallPopup();
  const selectedLink = useSelectedLink();

  callPopup("Delete", `Are you sure you want to delete ${selectedLink}?`, [
    {
      name: "Yes",
      callback: () => {
        console.log("Yes");
      },
    },
    {
      name: "No",
      callback: () => {
        console.log("No");
      },
    },
  ]);

  return <></>;
};
