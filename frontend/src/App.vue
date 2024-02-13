<template>
  <HelloWorld />
</template>

<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'

import { EventsOn } from '../wailsjs/runtime'

EventsOn('incoming', (data: any) => {
  console.log(
    `incoming message received with data ${JSON.stringify(data, null, 2)}`
  )
})
</script>
