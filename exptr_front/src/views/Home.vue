<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'

import usePage from '@/composables/page'

import { useOperationsStore } from '@/stores/operations'
import { useCategoriesStore } from '@/stores/categories'

import { TABLE_COLUMNS } from '@/helpers/consts'
import { PageAction } from '@/helpers/types'

import CategoryBtn from '@/components/categories/CategoryBtn.vue'

const { actionHandler } = usePage()

// DATA
const operationsStore = useOperationsStore()
const categoriesStore = useCategoriesStore()
onMounted(async () => {
  await Promise.all([operationsStore.fetchOperations(), categoriesStore.fetchCategories()])
})

// TABLE
const pagination = ref({
  sortBy: 'desc',
  descending: false,
  page: 2,
  rowsPerPage: 3
})
const pagesNumber = computed(() =>
  Math.ceil(operationsStore.operations.length / pagination.value.rowsPerPage)
)
</script>

<template>
  <div class="dashboard full-height column justify-between q-pa-sm">
    <div class="col">
      <div class="row items-start gap-16">
        <CategoryBtn
          v-for="category in categoriesStore.categories"
          :key="category.id"
          :category="category"
          @click="actionHandler(PageAction.TOGGLE_CATEGORY, category.id)"
        />
        <q-btn round flat icon="add" color="white" />
      </div>
    </div>
    <div class="col-7">
      <q-table
        v-model:pagination="pagination"
        flat
        bordered
        hide-pagination
        row-key="name"
        card-class="bg-transparent text-white"
        table-class="dashboard__operations-table"
        :rows="operationsStore.operations"
        :columns="TABLE_COLUMNS"
      />
      <div class="row justify-center q-mt-sm">
        <q-pagination v-model="pagination.page" color="warning" :max="pagesNumber" size="sm" />
      </div>
    </div>
  </div>
</template>

<style>
.q-table__container {
  height: 95%;
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

.q-field__native {
  color: white;
}

.q-field__append {
  color: white;
}
</style>
