<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { ref } from 'vue'
import 'webauthn-components/registration'

const email = ref('')
const password = ref('')

const userStore = useUserStore()

const onSubmit = async () => {
  if (window.PublicKeyCredential) {
    const available =
      await window.PublicKeyCredential.isUserVerifyingPlatformAuthenticatorAvailable()
    console.log(available)
    userStore.webAuthnSignUp(email.value)
  }
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
        outlined
        v-model="email"
        label="Your email *"
        lazy-rules
        color="warning"
        label-color="warning"
        input-style="color: orange"
        :rules="[(val) => (val && val.length > 0) || 'Please type something']"
      />

      <div>
        <q-btn label="Submit" type="submit" color="warning" />
        <q-btn label="Reset" type="reset" color="warning" flat class="q-ml-sm" />
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
