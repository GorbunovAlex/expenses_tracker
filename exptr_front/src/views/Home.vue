<script setup lang="ts">
import BarsChart from '@/components/analytics/BarsChart.vue'
import DoughtChart from '@/components/analytics/DoughtnutChart.vue'
</script>

<template>
  <div class="home q-pa-md">
    <div class="">
      <BarsChart />
    </div>
    <div class="">
      <DoughtChart />
    </div>
    <div class="row items-center justify-end text-white text-bold text-h5">$: 1234</div>
  </div>
</template>

<style scoped lang="scss">
.home {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
}
</style>
