import { atom, useAtomValue, useSetAtom } from "jotai";

type PopupAction = {
  name: string;
  callback: () => void;
};

type PopupAtom = {
  title: string;
  content: string;
  open: boolean;
  actions: PopupAction[];
};

const INITIAL_POPUP: PopupAtom = {
  title: "",
  content: "",
  open: false,
  actions: [],
};

const popupAtom = atom<PopupAtom>(INITIAL_POPUP);

export const usePopupValues = () => {
  return useAtomValue(popupAtom);
};

export const useCallPopup = () => {
  const set = useSetAtom(popupAtom);
  const setPopupInfos = (
    title: string,
    content: string,
    actions: PopupAction[]
  ) => {
    const actionsWithClose: PopupAction[] = actions.map((action) => ({
      ...action,
      callback: () => {
        action.callback();
        set(INITIAL_POPUP);
      },
    }));
    set({
      title,
      content,
      open: true,
      actions: actionsWithClose,
    });
  };

  return setPopupInfos;
};
