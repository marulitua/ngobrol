<template lang="pug">
  .ajax(v-if="isShown === true")
    control-icon(:icon="'ajax'", :classes="'fa-spin fa-3x fa-fw'")
</template>

<script>
import ControlIcon from './ControlIcon.vue'

export default {
  components: {
    'control-icon': ControlIcon
  },
  'data': function() {
    return {
      isShown: true,
    };
  },
  methods: {
    hide: function() {
      this.isShown = false;
    },
    show: function() {
      this.isShown = true;
    }
  }
}
</script>
