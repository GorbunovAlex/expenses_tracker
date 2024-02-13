import axios, { type AxiosInstance } from 'axios';

import { RequestMethod } from '@/helpers/types';
import ErrorHandler from '@/helpers/funcs/error-handler';
// import { getToken } from "@/helpers/funcs/auth-utils";

class ApiService {
  private static instance: ApiService;
  private axiosInstance: AxiosInstance;

  private constructor() {
    this.axiosInstance = axios.create({
      baseURL: process.env.API_PATH ?? 'http://localhost:3000',
      headers: {
        'Content-Type': 'application/json',
      },
    });
  }

  public static getInstance(): ApiService {
    if (!ApiService.instance) {
      ApiService.instance = new ApiService();
    }

    return ApiService.instance;
  }

  public static async request<T>(type: RequestMethod, url: string, data: T) {
    try {
      return await ApiService.instance.axiosInstance.request({
        method: type,
        url,
        data,
      });
    } catch (error) {
      new ErrorHandler(error as Error).handleError()
      return error;
    }
  }
}

export default ApiService.getInstance();