<script setup lang="ts">
import { useRoute } from 'vue-router'
import { date } from 'quasar'

import { PAGE_LINKS, PAGE_ACTIONS } from '@/helpers/consts'

const route = useRoute()
</script>

<template>
  <q-layout view="lHh lpR lFf" class="default-layout">
    <q-drawer show-if-above :width="50">
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
              <q-icon
                size="20px"
                :name="item.icon"
                :color="route.path === item.to ? 'warning' : 'white'"
              />
            </q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>

    <q-page-container class="flex items-center justify-center">
      <q-page class="page q-pa-md">
        <div class="flex items-center justify-between text-h5 text-white">
          {{ route.name }}
          <div class="flex items-center gap-8 text-white">
            <div class="text-white text-body1">$: 1234</div>
            <q-separator vertical size="2px" color="white" />
            <span class="text-body1">
              {{ date.formatDate(new Date(), 'DD/MM/YYYY') }}
            </span>
            <q-separator vertical size="2px" color="white" />
            <q-btn-dropdown flat rounded dense size="15px" dropdown-icon="settings">
              <q-list class="page__settings-list">
                <q-item v-for="item in PAGE_ACTIONS" :key="item.label" clickable>
                  <q-item-section>{{ item.label }}</q-item-section>
                </q-item>
              </q-list>
            </q-btn-dropdown>
          </div>
        </div>
        <q-separator class="q-ma-sm" color="white" />
        <div class="page__content">
          <router-view />
        </div>
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
  background: rgba(242, 192, 55, 0.2);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(7.5px);
  -webkit-backdrop-filter: blur(7.5px);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.18);
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
  width: 80%;
  background: rgba(242, 192, 55, 0.4);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.18);
  border: 1px solid $warning;
  border-radius: 10px;

  &__settings-list {
    background: rgba(242, 192, 55, 0.4);
    box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
  }

  &__content {
    height: 75vh;
  }
}
</style>
