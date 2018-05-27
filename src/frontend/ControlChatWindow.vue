<template lang="pug">
  .chat-window(:class="{'expanded': isExpanded}", :ref="'index-' + index")
    .header(@click="expand")
      .name {{ usersInChat }}
      .close(@click="closeChat")
        control-icon.icon(:icon="'close'")
    .body(:ref="'chat'")
      ul.chat
        li.line(v-if="chatdata", v-for="(hist, index) in chatdata.history", :class="{'me': parseInt(hist.user) === 0, 'them': parseInt(hist.user) !== 0 }")
          .bubble
            //- .name {{ getUser(hist.user) }}
            .message {{ hist.message }}
          .time
            | {{ time }}
        //- li.line(v-if="iAmTyping", class="me")
        //-   .bubble
        //-     .message
        //-       .dot-1
        //-       .dot-2
        //-       .dot-3
        li.line(v-if="theyAreTyping", class="them")
          .bubble
            .message
              .dot-1
              .dot-2
              .dot-3
    .footer
      .chat-input
        input(:ref="'send-message'", v-model="sendMessageInput", @keyup.enter="sendMessage", placeholder="type something...")
        button.wobbler(@click="$root.$children[0].wobble")
          control-icon.icon(:icon="'emoji'")
        button.wobbler(@click="$root.$children[0].wobble")
          control-icon.icon(:icon="'attachment'")
        button.wobbler(@click="sendMessage", :class="{'disabled': sendMessageInput === ''}")
          control-icon.icon(:icon="'send'")
</template>

<script>
import ControlIcon from './ControlIcon.vue'
import './tweenjs.js'

export default {
  components: {
    'control-icon': ControlIcon
  },
  data: function() {
    return {
      isExpanded: true,
      iAmTyping: false,
      theyAreTyping: false,
      sendMessageInput: '',
      readTimeout: false,
      writeTimeout: false,
      scrollbar: false,
    };
  },
  props: {
    index: {
      type: Number,
      required: true,
    },
    chatdata: {
      type: Object,
      required: true
    },
  },
  computed: {
    time: function() {
      let hour = Math.ceil(Math.random() * 12);
      let minute = Math.ceil(Math.random() * 60);
      let ampm = Math.ceil(Math.random() * 2) === 2? 'pm' : 'am';
      return `${hour}:${minute > 9 ? minute : '0' + minute} ${ampm}`;
    },
    usersInChat: function() {
      let userList = this.getUsers();
      userList.shift();
      return userList.join(', ');
    },
  },
  methods: {
    closeChat: function(ev) {
      ev.preventDefault();
      ev.stopPropagation();
      this.$root.$children[0].removeChat(this.index)
    },
    expand: function() {
      this.isExpanded = !this.isExpanded;
      this.scrollToBottom();
    },
    getUser: function(userIndex) {
      return this.chatdata.users[userIndex].username;
    },
    getUsers: function() {
      let users = [];
      this.chatdata.users.forEach((user) => {
        users.push(user.username);
      });

      return users;
    },
    scrollToBottom: function(force) {
      force = force || false;
      let chat = this.$refs.chat;
      if (!force) {
        var obj = createjs.Tween.get(chat, {
        }).to({
          scrollTop: chat.scrollHeight
        }, 600).play();
      } else {
        chat.scrollTop = chat.scrollHeight;
      }
    },
    sendMessage: function(ev) {
      let messages = [
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc venenatis augue accumsan.',
        'Cras ullamcorper turpis tellus, sit amet tristique turpis fermentum eget. Vivamus iaculis.',
        'Nulla et massa feugiat, lobortis quam vitae, euismod urna. Vestibulum blandit ultrices.',
        'Morbi laoreet sollicitudin auctor. Nunc eu augue blandit, vulputate orci at, maximus.',
        'Nullam nec auctor massa. Sed felis ipsum, tincidunt vitae dapibus eu, ornare.'
      ]
      this.chatdata.history.push({
        user: 0,
        message: this.sendMessageInput
      })
      this.sendMessageInput = '';
      this.scrollToBottom();
      let readTime = Math.floor(Math.random() * 5);
      if (this.readTimeout) {
        clearTimeout(this.readTimeout);
      }
      if (this.writeTimeout) {
        clearTimeout(this.writeTimeout);
      }
      this.readTimeout = setTimeout(() => {
        let writeTime = Math.floor(Math.random() * 8);
        this.theyAreTyping = true;
        this.scrollToBottom();

        if (this.writeTimeout) {
          clearTimeout(this.writeTimeout);
        }
        this.writeTimeout = setTimeout(() => {
          let messageIndex = Math.floor(Math.random() * 5);
          let message = messages[messageIndex];

          this.chatdata.history.push({
            user: 1,
            message: message
          })
          this.theyAreTyping = false;
          this.scrollToBottom();
        }, writeTime * 1000);
      }, readTime * 1000);
    },
  },
  watch: {
    sendMessageInput: function(value) {
      if (value !== '') {
        this.iAmTyping = true;
      } else {
        this.iAmTyping = false;
      }
      this.scrollToBottom();
    }
  },
  mounted: function() {
    this.scrollToBottom(true);
    let sendMessage = this.$refs['send-message'];
    if (sendMessage) {
      sendMessage.focus();
    }
    // this.scrollbar = new PerfectScrollbar(this.$refs.chat);
    // console.log(this.scrollbar)
  },
};
</script>
