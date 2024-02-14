<script setup lang="ts">
import { onMounted } from 'vue'
import { RouterView } from 'vue-router'

import { useUserStore } from '@/stores/user'

const userStore = useUserStore()

const mockAuth = () => {
  const name = prompt('email')
  const pass = prompt('password')
  if (!name || !pass) {
    mockAuth()
  } else {
    userStore.login(name, pass)
  }
}

onMounted(() => {
  mockAuth()
})
</script>

<template>
  <RouterView />
</template>
