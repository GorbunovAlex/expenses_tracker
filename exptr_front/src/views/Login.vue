<script setup lang="ts">
import { ref } from 'vue'

const email = ref('')
const password = ref('')

const onSubmit = () => {
  console.log('Submitted!', email.value, password.value)
}

const onReset = () => {
  email.value = ''
  password.value = ''
}
</script>

<template>
  <div class="login">
    <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
      <q-input
        filled
        v-model="email"
        label="Your email *"
        lazy-rules
        :rules="[(val) => (val && val.length > 0) || 'Please type something']"
      />

      <q-input
        filled
        v-model="password"
        label="Password *"
        lazy-rules
        :rules="[(val) => (val !== null && val !== '') || 'Please type your age']"
      />

      <div>
        <q-btn label="Submit" type="submit" color="primary" />
        <q-btn label="Reset" type="reset" color="primary" flat class="q-ml-sm" />
      </div>
    </q-form>
  </div>
</template>

<style scoped lang="scss">
.login {
  width: 30%;
  height: 40%;
  display: grid;
  place-items: center;
  padding: 10px;
  background: rgba(242, 192, 55, 0.4);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.18);
  border: 1px solid $warning;
  border-radius: 10px;
}
</style>
