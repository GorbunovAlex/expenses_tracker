<script setup lang='ts'>
import { PAGE_LINKS } from '@/helpers/consts';
</script>

<template>
<q-layout view="lHh lpR lFf" class="default-layout">
    <q-drawer
      show-if-above
      :width="50"
    >
      <q-list class="drawer-list">
        <q-item
          v-for="item in PAGE_LINKS"
          :key="item.icon"
          clickable
          :to="item.to"
          active-class="text-warning"
        >
          <q-item-section>
            <q-item-label>
              <q-icon size="20px" :name="item.icon" />
            </q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>

    <q-page-container class="flex items-center justify-center">
      <q-page class="page">
        <router-view />
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<style lang="scss">
.default-layout {
  position: relative;
}

.q-drawer {
  height: 300px;
  position: fixed;
  background: rgba( 242, 192, 55, 0.2 );
  box-shadow: 0 8px 32px 0 rgba( 31, 38, 135, 0.37 );
  backdrop-filter: blur( 7.5px );
  -webkit-backdrop-filter: blur( 7.5px );
  border-radius: 10px;
  border: 1px solid rgba( 255, 255, 255, 0.18 );
  top: 35%;
  left: 30px;
  border-radius: 10px;
  border: 1px solid $warning;
}

.drawer-list {
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  height: 100%;
}

.q-page-container {
  height: 100vh;
}

.page {
  min-height: 80vh !important;
  height: 50%;
  width: 80%;
  padding: 20px;
  background: rgba( 242, 192, 55, 0.4 );
  box-shadow: 0 8px 32px 0 rgba( 31, 38, 135, 0.37 );
  backdrop-filter: blur( 10px );
  -webkit-backdrop-filter: blur( 10px );
  border-radius: 10px;
  border: 1px solid rgba( 255, 255, 255, 0.18 );
  border: 1px solid $warning;
  border-radius: 10px;
}
</style>