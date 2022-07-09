const API_LINK = "http://localhost:3001/";

export const apiRoutes = {
  getLinks: (email: string) => `${API_LINK}get?email=${email}`,
  createLink: `${API_LINK}create`,
  deleteLink: (id: string) => `${API_LINK}delete?id=${id}`,
  editLink: `${API_LINK}edit`,
  getAnalytics: (id: string) => `${API_LINK}analytics?id=${id}`,
};
