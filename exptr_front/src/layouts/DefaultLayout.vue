<script setup lang='ts'></script>

<template>
<q-layout view="lHh lpR lFf" class="bg-white">
    <q-drawer
      show-if-above
      bordered
      class="bg-grey-2"
    >
      
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>