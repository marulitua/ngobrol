<template lang="pug">
  li.chat-user(@click="chatUser(index)", :class="{'offline': userdata.status === 0, 'online': userdata.status === 1, 'away': userdata.status === 2, }")
    .photo-container
      .photo
        img(:src="userdata.photo")
    .username {{ userdata.username}}
    .last-message
      .sender You:&nbsp;
      | lorem ipsum dolor sit amet
    .notifications(:class="{'show': Math.ceil(Math.random() * 10) % 2 }")
      .text {{ Math.ceil(Math.random() * 10) }}
</template>

<script>
export default {
  data: function() {
    return {
    };
  },
  props: {
    index: {
      type: Number,
      required: true,
    },
    userdata: {
      type: Object,
      required: true,
    },
  },
  computed: {
  },
  methods: {
    chatUser: function(index) {
      let $root = this.$root.$children[0]
      if ($root.lookupChatByIndex(index) === false) { // check to see if the chat is already opened
        $root.chatWith(index, this.userdata);
      } else {
        $root.focusChat(index)
      }
    },
  },
  watch: {
  },
  mounted: function() {
  }
}
</script>
