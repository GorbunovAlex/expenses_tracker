<script setup lang="ts">
import { onMounted } from 'vue'

import { useOperationsStore } from '@/stores/operations'
import { useCategoriesStore } from '@/stores/categories'

import { TABLE_COLUMNS } from '@/helpers/consts'

import CategoryBtn from '@/components/dashboard/CategoryBtn.vue'

const operationsStore = useOperationsStore()
const categoriesStore = useCategoriesStore()

onMounted(async () => {
  await Promise.all([operationsStore.fetchOperations(), categoriesStore.fetchCategories()])
  console.log('Operations:', operationsStore.operations)
  console.log('Categories:', categoriesStore.categories)
})
</script>

<template>
  <div class="dashboard full-height column justify-between q-pa-sm">
    <div class="col">
      <div class="row items-start gap-16">
        <CategoryBtn
          v-for="category in categoriesStore.categories"
          :key="category.id"
          :category="category"
        />
        <q-btn round flat icon="add" color="white" />r
      </div>
    </div>
    <div class="col-7">
      <q-table
        flat
        bordered
        card-class="bg-transparent text-white"
        table-class="dashboard__operations-table"
        :rows="operationsStore.operations"
        :columns="TABLE_COLUMNS"
        row-key="name"
      />
    </div>
  </div>
</template>

<style>
.q-table__container {
  height: 100%;
}

.q-table--bordered {
  border: 1px solid white;
}

.q-table thead,
.q-table tr,
.q-table th,
.q-table td {
  border-color: white;
}

.q-table__bottom {
  border-top: 1px solid white;
}
</style>
