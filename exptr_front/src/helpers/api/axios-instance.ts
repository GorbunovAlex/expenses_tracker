import Axios, { AxiosError, type AxiosRequestConfig } from 'axios';

import ErrorHandler from '@/helpers/funcs/error-handler';
import { getToken } from '../funcs/auth-utils';
 
 export const AXIOS_INSTANCE = Axios.create({ baseURL: '/api/v1/' }); 

 export const instance = <T>(
  config: AxiosRequestConfig,
  options?: AxiosRequestConfig,
): Promise<T> => {
  const source = Axios.CancelToken.source();
  if (config.url && !config.url.includes('users')) {
    config.headers = {
      "Bearer": getToken(),
    };
  }
  const promise = AXIOS_INSTANCE({
    ...config,
    ...options,
    cancelToken: source.token,
  })
    .then(({ data }) => data)
    .catch((error) => {
      new ErrorHandler(error as Error).handleError()
      throw error;
    });

  // @ts-ignore
  promise.cancel = () => {
    source.cancel('Query was cancelled');
  };

  return promise;
};

export type ErrorType<Error> = AxiosError<Error>;