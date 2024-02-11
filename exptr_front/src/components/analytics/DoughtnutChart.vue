<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Chart, registerables } from 'chart.js'

Chart.register(...registerables)

const doughnut = ref<HTMLCanvasElement | null>(null)
const chartData = ref({
  labels: ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'],
  datasets: [
    {
      label: '# of Votes',
      data: [12, 19, 3, 5, 2, 3],
      backgroundColor: [
        'rgba(255, 99, 132, 0.2)',
        'rgba(54, 162, 235, 0.2)',
        'rgba(255, 206, 86, 0.2)',
        'rgba(75, 192, 192, 0.2)',
        'rgba(153, 102, 255, 0.2)',
        'rgba(255, 159, 64, 0.2)'
      ],
      borderColor: [
        'rgba(255, 99, 132, 1)',
        'rgba(54, 162, 235, 1)',
        'rgba(255, 206, 86, 1)',
        'rgba(75, 192, 192, 1)',
        'rgba(153, 102, 255, 1)',
        'rgba(255, 159, 64, 1)'
      ],
      borderWidth: 2
    }
  ]
})
const chartOptions = ref({
  scales: {
    y: {
      beginAtZero: true
    }
  }
})

onMounted(() => {
  if (doughnut.value) {
    const ctx = doughnut.value.getContext('2d')
    if (ctx) {
      new Chart(ctx, {
        type: 'doughnut',
        data: chartData.value,
        options: chartOptions.value
      })
    }
  }
})
</script>

<template>
  <canvas ref="doughnut" />
</template>
