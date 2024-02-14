<script setup lang="ts">
import { onMounted } from 'vue'

import { useOperationsStore } from '@/stores/operations'
import { useCategoriesStore } from '@/stores/categories'

import { TABLE_COLUMNS } from '@/helpers/consts'
import DoughtChart from '@/components/analytics/DoughtnutChart.vue'

const operationsStore = useOperationsStore()
const categoriesStore = useCategoriesStore()

const rows = [
  {
    name: 'Frozen Yogurt',
    calories: 159,
    fat: 6.0,
    carbs: 24,
    protein: 4.0,
    sodium: 87,
    calcium: '14%',
    iron: '1%'
  }
]

onMounted(async () => {
  await Promise.all([operationsStore.fetchOperations(), categoriesStore.fetchCategories()])
  console.log('operationsStore.operations', operationsStore.operations)
  console.log('categoriesStore.categories', categoriesStore.categories)
})
</script>

<template>
  <div class="row gap-16 q-pa-md">
    <div class="col">
      <DoughtChart />
    </div>
    <div class="col">
      <div class="col text-white text-bold text-h5">$: 1234</div>
      <q-table title="Treats" :rows="rows" :columns="TABLE_COLUMNS" row-key="name" />
    </div>
  </div>
</template>
