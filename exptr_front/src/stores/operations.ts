import { reactive } from 'vue'
import { defineStore } from 'pinia'

import type { ModelsOperation } from '@/helpers/api/models'
import { getOperations } from '@/helpers/api/api'

export const useOperationsStore = defineStore('operations', () => {
  const operations = reactive<ModelsOperation[]>([])

  async function fetchOperations() {
    try {
      const response = await getOperations()
      response.operations?.forEach((opeartion) => {
        operations.push(opeartion)
      });
    } catch (error) {
      console.error(error)
    }
  }
 
  return {
    operations,

    fetchOperations
  }
})
