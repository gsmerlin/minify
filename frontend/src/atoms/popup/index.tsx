import { atom, useAtomValue, useSetAtom } from "jotai";
import React from "react";

type PopupAtom = {
  content: React.ReactNode;
  open: boolean;
};

const INITIAL_POPUP = {
  content: null,
  open: false,
};

const popupAtom = atom<PopupAtom>(INITIAL_POPUP);

export const usePopupValues = () => {
  return useAtomValue(popupAtom);
};

interface PopupClose {
  close: true;
  content?: never;
}
interface PopupContent {
  close?: never;
  content: React.ReactNode;
}

type CallPopup = PopupContent | PopupClose;

export const useCallPopup = () => {
  const set = useSetAtom(popupAtom);
  const setPopupInfos = (params: CallPopup) => {
    if (params.close) {
      set({ ...INITIAL_POPUP });
      return;
    }
    if (params.content) {
      set({ ...params, open: true });
    }
  };

  return setPopupInfos;
};
