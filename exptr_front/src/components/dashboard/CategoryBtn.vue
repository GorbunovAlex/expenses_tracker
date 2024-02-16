<script setup lang="ts">
import type { ModelsCategory } from '@/helpers/api/models'

defineProps<{
  category: ModelsCategory
}>()
</script>

<template>
  <div class="column items-center gap-4">
    <q-btn :id="`category` + category.id" round color="warning" />
    <label :for="`category` + category.id" class="text-white"> {{ category.name }} </label>
  </div>
</template>
