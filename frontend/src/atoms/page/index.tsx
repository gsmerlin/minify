import { atom, useAtomValue, useSetAtom } from "jotai";
// This is more for playing around,
// you shouldn't really do routing with atoms
// but I'm curious how difficult would it be to
// implement a simple routing solution.

export enum Page {
  Login = "login",
  Home = "home",
  Edit = "edit",
  Delete = "delete",
  Analytics = "analytics",
}

const pageAtom = atom<string>(Page.Login);

export const usePage = () => {
  return useAtomValue(pageAtom);
};

export const usePageActions = () => {
  const setPageAtom = useSetAtom(pageAtom);
  return {
    setPage: (newPage: Page) => setPageAtom(newPage),
  };
};
