const checkIsURL = (url: string) => {
  const pattern = /^(http|https):\/\/[^ "]+$/;
  return pattern.test(url);
};

const checkIsExistingURL = async (url: string) => {
  try {
    await fetch(url, {
      mode: "no-cors",
    });
    return true;
  } catch {
    return false;
  }
};

export const validateURL = async (url: string) => {
  const isURL = checkIsURL(url);
  const isExistingURL = await checkIsExistingURL(url);
  if (!isURL || !isExistingURL) {
    return "Please enter a valid URL!";
  }
};