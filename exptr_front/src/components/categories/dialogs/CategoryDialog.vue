<script setup lang="ts">
import { useDialogPluginComponent } from 'quasar'

defineEmits([...useDialogPluginComponent.emits])

defineProps<{
  id: string
}>()

const { dialogRef, onDialogHide, onDialogOK, onDialogCancel } = useDialogPluginComponent()
</script>

<template>
  <q-dialog ref="dialogRef" @hide="onDialogHide">
    <q-card>
      <q-card-section>
        <div class="text-h6">Dialog Title</div>
      </q-card-section>
      <q-card-section> CONTENT </q-card-section>
      <q-card-actions>
        <q-btn color="primary" label="OK" @click="onDialogOK" />
        <q-btn color="primary" label="Cancel" @click="onDialogCancel" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>
