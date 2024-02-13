import Cookies from 'js-cookie';

export const getToken = (): string | undefined => Cookies.get('sid');

export const setToken = (token?: string | null): undefined => {
  if (!token) return;

  Cookies.set('sid', token);
};

export const clearToken = (): void => Cookies.remove('sid');