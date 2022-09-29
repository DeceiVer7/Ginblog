<template>
  <div id="particles"></div>
</template>

<script>
import particlesJson from './particles.json'
export default {
  data () {
    return {}
  },
  mounted () {
    require('./particles')
    // eslint-disable-next-line no-undef
    particlesJS('particles', particlesJson)
  },
  name: 'particles'
}
</script>

<style scoped>
#particles {
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: #003472;
  background-repeat: no-repeat;
  background-size: cover;
  background-position: 50% 50%;
}
</style>
