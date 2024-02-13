<template>
  <main>
    <div>Hello, world!</div>

    <button class="btn btn-primary" @click="dingBell">Ring the bell</button>
  </main>
</template>

<script lang="ts" setup>
import { Greet } from '../../wailsjs/go/main/App'

const dingBell = async () => {
  console.log(`dingBell called`)
  const ans = await Greet('foobar')
  console.log(`... Greet returns ${JSON.stringify(ans)}`)
}
</script>
