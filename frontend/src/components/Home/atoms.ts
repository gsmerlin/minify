import axios from 'axios';
import { atom, useSetAtom, useAtomValue } from 'jotai'
import { emailAtom } from './../../atoms/user/index';

export type Link = {
    ID: string;
    Email: string;
    Destination: string;
}

export type Links = Link[]

const fetchLinksAtom = atom<Promise<Links>>(async (get) => {
    const email = get(emailAtom);
    const { data } = await axios.get<Links>(`http://localhost:3001/get?email=${email}`);
    return data;
});

const linksAtom = atom<Links>([])

const selectedLinkAtom = atom<string>('')

export const useFetchLinks = () => {
    const links = useAtomValue(fetchLinksAtom);
    // Sets the local links atom
    useSetAtom(linksAtom)(links);
    return links
}

export const useGetLocalLinks = () => {
    return useAtomValue(linksAtom)
}

export const useSelectedLink = () => {
    const links = useGetLocalLinks();
    const id = useAtomValue(selectedLinkAtom);
    return links.find(link => link.ID === id)
}

export const useLinkActions = () => {
    const setSelectedLink = useSetAtom(selectedLinkAtom)
    return {
        setSelectedLink: (id: string) => setSelectedLink(id)
    }
}