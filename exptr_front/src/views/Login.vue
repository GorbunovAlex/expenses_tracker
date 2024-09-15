<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import { useUserStore } from '@/stores/user'

const email = ref('')
const password = ref('')
const userStore = useUserStore()
const router = useRouter();
async function onSubmit() {
  if (email.value && password.value) {
    const result = await userStore.login(email.value, password.value)
    if (result) {
      router.push({ name: 'Dashboard' })
    }
  }
}

function onReset() {
  email.value = ''
  password.value = ''
}
</script>

<template>
  <div class="login">
    <q-form @submit="onSubmit" @reset="onReset" class="login__form">
      <div class="text-h6 text-white q-mb-md">Login</div>
      <q-input
        outlined
        v-model="email"
        label="Email *"
        lazy-rules
        color="warning"
        label-color="white"
        input-style="color: white"
        :rules="['email']"
      />
      <q-input
        outlined
        v-model="password"
        label="Password *"
        lazy-rules
        color="warning"
        type="password"
        label-color="white"
        input-style="color: white"
        :rules="[(val) => (val && val.length > 0) || 'Password is required']"
      />

      <div class="flex row justify-between items-center">
        <div class="text-body1 text-white">Don't have an account? <q-btn dense flat class="text-warning">Sign up</q-btn></div>
        <div class="flex row">
          <q-btn label="Submit" type="submit" color="warning" />
          <q-btn label="Reset" type="reset" color="warning" flat class="q-ml-sm" />
        </div>
      </div>
    </q-form>
  </div>
</template>

<style>
.q-field--outlined .q-field__control:before {
  border: 1px solid rgb(183, 159, 113);
}
</style>

<style scoped lang="scss">
.login {
  width: 60%;
  display: grid;
  place-items: center;
  padding: 10px;
  background: rgba(242, 192, 55, 0.4);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.18);

  &__form {
    width: 100%;
  }
}
</style>
