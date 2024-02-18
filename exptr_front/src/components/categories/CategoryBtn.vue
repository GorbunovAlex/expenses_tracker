<script setup lang="ts">
import type { ModelsCategory } from '@/helpers/api/models'

defineProps<{
  category: ModelsCategory
}>()
const emit = defineEmits(['click'])
</script>

<template>
  <div class="column items-center gap-4">
    <q-btn
      :id="`category` + category.id"
      :icon="category.icon === '' ? 'home' : category.icon"
      :color="category.color === '' ? 'warning' : category.color"
      round
      @click="emit('click')"
    />
    <label :for="`category` + category.id" class="text-white"> {{ category.name }} </label>
  </div>
</template>
