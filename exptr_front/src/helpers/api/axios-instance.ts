import Axios, { AxiosError, type AxiosRequestConfig } from 'axios';

import ErrorHandler from '@/helpers/funcs/error-handler';
 
 export const AXIOS_INSTANCE = Axios.create({ baseURL: import.meta.env.API_PATH ?? 'http://localhost:3000/v1/api/' }); 

 export const instance = <T>(
  config: AxiosRequestConfig,
  options?: AxiosRequestConfig,
): Promise<T> => {
  const source = Axios.CancelToken.source();
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