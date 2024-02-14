import { reactive } from 'vue'
import { defineStore } from 'pinia'

import type { ModelsCategory } from '@/helpers/api/models'
import { getCategories } from '@/helpers/api/api'

export const useCategoriesStore = defineStore('categories', () => {
  const categories = reactive<ModelsCategory[]>([])

  async function fetchCategories() {
    try {
      const response = await getCategories()
      response.categories?.forEach((category) => {
        categories.push(category)
      });
    } catch (error) {
      console.error(error)
    }
  }
 
  return {
    categories,

    fetchCategories
  }
})
