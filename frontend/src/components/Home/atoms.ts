import axios from 'axios';
import { atom, useAtomValue } from 'jotai'
import { emailAtom } from './../../atoms/user/index';

export type Links = {
    ID: string;
    Email: string;
    Destination: string;
}[]

const linksAtom = atom<Promise<Links>>(async (get) => {
    const email = get(emailAtom);
    const { data } = await axios.get<Links>(`http://localhost:3001/get?email=${email}`);
    return data;
});

export const useLinks = () => {
    const links = useAtomValue(linksAtom);
    return links
}