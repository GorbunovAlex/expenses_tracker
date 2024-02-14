import { defineStore } from 'pinia'

import { postUsersLogin } from '@/helpers/api/api'
import { setToken } from '@/helpers/funcs/auth-utils'

export const useUserStore = defineStore('user', () => {
  async function login(email: string, password: string) {
    try {
      const response = await postUsersLogin({ email, password })
      if (response.status === "ok") {
        setToken(response.token);
      }
    } catch (error) {
      console.error(error)
    }
  }

  return {
    login
  }
})
