import axios from "axios";
import { atom, useSetAtom, useAtomValue } from "jotai";
import { apiRoutes } from "../../globals";
import { emailAtom } from "../../atoms/user/index";

export type Link = {
  ID: string;
  Email: string;
  Destination: string;
};

export type Links = Link[];

const linksAtom = atom<Links>([]);

const fetchLinksAtom = atom(null, async (get, set) => {
  const email = get(emailAtom);
  const { data } = await axios.get<Links>(apiRoutes.getLinks(email));
  set(linksAtom, data);
});

export const useLinks = () => {
  const fetchLinks = useSetAtom(fetchLinksAtom);
  // onMount only runs when Atom is first used.
  // Therefore this allows us to fetch the atom values
  // on first load
  if (!linksAtom.onMount) {
    linksAtom.onMount = () => {
      fetchLinks();
    };
  }

  const links = useAtomValue(linksAtom);

  return links;
};

const selectedLinkAtom = atom<string>("");

// Here we use an atom so we can create a link based on the user email
// without subscribing to the email atom in the main hook (thus preventing)
// re-rendering cycles.
// Since setter atoms can't return values, this does force us to create a new
// atom to store it's return value, but we can abstract that into our hook helper.
const createdLinkAtom = atom<Link>({ ID: "", Email: "", Destination: "" });
const createLinkAtom = atom(null, async (get, set, destination: string) => {
  const email = get(emailAtom);
  const { data } = await axios.post<Link>(apiRoutes.createLink, {
    email,
    destination,
  });
  set(createdLinkAtom, data);
});

type LinkAnalytics = {
  ID: string;
  Email: string;
  Destination: string;
  TotalClicks: number;
  Timestamps: string[];
};

const linkAnalyticsAtom = atom<LinkAnalytics>({
  ID: "",
  Email: "",
  Destination: "",
  TotalClicks: 0,
  Timestamps: [],
});

type LinkAnalyticsJson = {
  id: string;
  timestamps: string[];
};

const getLinkAnalyticsAtom = atom(null, async (get, set) => {
  const selectedLink = get(selectedLinkAtom);
  const links = get(linksAtom);
  const link = links.find((l) => l.ID === selectedLink);
  const { data } = await axios.get<LinkAnalyticsJson>(
    apiRoutes.getAnalytics(selectedLink)
  );
  set(linkAnalyticsAtom, {
    ID: link?.ID ?? "",
    Email: link?.Email ?? "",
    Destination: link?.Destination ?? "",
    TotalClicks: data.timestamps ? data.timestamps.length : 0,
    Timestamps: data.timestamps ?? [],
  });
});

export const useLinkAnalytics = () => {
  useSetAtom(getLinkAnalyticsAtom)();

  const linkAnalytics = useAtomValue(linkAnalyticsAtom);

  return linkAnalytics;
};

export const useSelectedLink = () => {
  const links = useLinks();
  const id = useAtomValue(selectedLinkAtom);
  return links.find((link) => link.ID === id);
};

export const useGetCreatedLink = () => {
  return useAtomValue(createdLinkAtom);
};

export const useLinkActions = () => {
  const setSelectedLink = useSetAtom(selectedLinkAtom);
  const fetchLinks = useSetAtom(fetchLinksAtom);
  const createLink = useSetAtom(createLinkAtom);
  return {
    setSelectedLink: (id: string) => setSelectedLink(id),
    fetchLinks: () => fetchLinks(),
    deleteLink: async (id: string) => {
      const result = await axios.delete(apiRoutes.deleteLink(id));
      if (result.status === 200) {
        fetchLinks();
      }
      return result.status;
    },
    createLink: async (destination: string) => {
      await createLink(destination);
      fetchLinks();
    },
    editLink: async (link: Link) => {
      const result = await axios.put(apiRoutes.editLink, link);
      if (result.status === 200) {
        fetchLinks();
      }
      return result.status;
    },
  };
};
