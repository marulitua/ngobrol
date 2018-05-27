<template lang="pug">
  #app
    h1 Vue Chat Control
    h6
      b Note:
      span  This chat is not backed by any backend. Chats are simulated.
    p
      .developer
        | Developed By:
        a(href="https://codepen.io/phreaknation", target="_blank" title="") Joel Dies
        | &nbsp;|&nbsp;
        a(href="https://twitter.com/PhreakNation", target="_blank" title="") Twitter
        | &nbsp;|&nbsp;
        a(href="https://www.linkedin.com/in/joel-dies-30650025/", target="_blank" title="") LinkedIn
      .designer Designed by:&nbsp;
        a(href="https://codepen.io/rlignitz/", target="_blank" title="") Shea Lignitz
        | &nbsp;|&nbsp;
        a(href="https://twitter.com/shealignitz", target="_blank" title="") Twitter
        | &nbsp;|&nbsp;
        a(href="https://www.linkedin.com/in/rlignitz/", target="_blank" title="") LinkedIn
    p
      .option
        h4 Pick iconset
      .option
        label
          input#fontawesome(type="radio", name="iconset", v-model="iconset", value="fontawesome")
          |  Font Awesome
      .option
        label
          input#materialdesign(type="radio", name="iconset", v-model="iconset", value="materialdesign")
          |  Material Design
    .chat-box
      .chat-windows(:ref="'chat-windows'")
        chat-window(v-for="(chat, index) in chatsReversed", :key="index" :chatdata="chat", :index="index")
      chat-panel(:ref="'chat-panel'", :title="'Vue.JS Chat'")
</template>

<script>
import ChatPanel from './ControlChatPanel.vue'
import ChatWindow from './ControlChatWindow.vue'
import './app.scss'

let url = {
  users: 'https://jsonblob.com/api/jsonBlob/34f2ac63-44d2-11e8-adbe-d35ed1c5ed6f',
  photos: 'https://randomuser.me/api/?results=10&format=json',
}
export default {
  components: {
    'chat-panel': ChatPanel,
    'chat-window': ChatWindow
  },
  name: 'app',
  data () {
    return {
      chats: [],
      validTypes: [
        'application/json',
        'vnd.api+json'
      ],
      iconset: 'fontawesome',
      icons: [],
      }
    },
    computed: {
      chatsReversed() {
        return this.chats.slice().reverse();
      }
    },
    watch: {
      iconset: function(value) {
        this.icons.forEach((icon) => {
          icon.iconset = value
        })
      }
    },
    methods: {
      buildChat: function(index, userdata) {
        let messages = [
          'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce vitae arcu at arcu ultrices dignissim. Cras.',
          'Aliquam nec urna vitae leo lacinia. Duis lectus, efficitur sit amet scelerisque ac, consectetur.',
          'Curabitur maximus eleifend odio vitae commodo. Duis risus.',
          'Suspendisse ullamcorper cursus pulvinar. Sed pretium purus ut.'
        ]
        let data = {
          users: [
            { username: 'John Doe' },
          ],
          history: []
        };
        data.users.push({
          username: userdata.username,
        });
        messages.forEach((message) => {
          let user = (Math.random()>0.5)? 1 : 0;
          data.history.push({
            user: user,
            message: message
          });
        });
        data.id = index;
        return data;
      },
      chatWith: function(index, userdata) {
        let _chat = this.buildChat(index, userdata);
        this.chats.unshift(_chat);
      },
      focusChat: function(index) {
        let chat = this.$refs['index-' + index];
        if (chat) {
          chat.$refs['send-message'].focus();
        }
      },
      getData: function(url) {
        return fetch(url, { cors: false}).then((response) => {
          if (!response.ok) {
            throw Error(response.statusText + '. ' + urlUsers);
          }
          let contentType = response.headers.get('content-type');
          let valid = false;
          if (contentType) {
            this.validTypes.forEach((type) => {
              if (contentType.includes(type)) {
                valid = true;
              }
            })
            if (valid === true) {
              return response.json();
            }
          }
          throw new TypeError('Content returned from "' + (url) + '" is not valid json. \ncontentType: ' + contentType);
        }).catch((err) => {
          console.error(err)
        });
      },
      lookupChatByIndex: function(index) {
        let ret = false;
        this.chats.forEach((chat) => {
          if (chat.id === index) {
            ret = chat;
          }
        })
        return ret;
      },
      removeChat: function(index) {
        this.chats.splice(this.index, 1);
      },
      wobble: function(ev) {
        ev.preventDefault();
        ev.stopPropagation();
        let el = ev.target;
        if (!el.classList.contains('wobbler')) {
          el = el.parentNode;
        }
        el.classList.add('wobble');
        setTimeout(() => {
          el.classList.remove('wobble');
        }, 500)
      }
    },
    mounted: function() {
      let chatPanel = this.$refs['chat-panel'];
      chatPanel.isLoaded = false;

      let userData = [  {    "id": 1,    "name": "Leanne Graham",    "status": 1,    "username": "Bret"  },  {    "status": 0,    "username": "Antonette",    "name": "Ervin Howell",    "id": 2  },  {    "status": 2,    "username": "Samantha",    "name": "Clementine Bauch",    "id": 3  },  {    "status": 2,    "username": "Karianne",    "name": "Patricia Lebsack",    "id": 4  },  {    "status": 1,    "username": "Kamren",    "name": "Chelsey Dietrich",    "id": 5  },  {    "status": 0,    "username": "Leopoldo_Corkery",    "name": "Mrs. Dennis Schulist",    "id": 6  },  {    "status": 1,    "username": "Elwyn.Skiles",    "name": "Kurtis Weissnat",    "id": 7  },  {    "status": 0,    "username": "Maxime_Nienow",    "name": "Nicholas Runolfsdottir V",    "id": 8  },  {    "status": 1,    "username": "Delphine",    "name": "Glenna Reichert",    "id": 9  },  {    "status": 0,    "username": "Moriah.Stanton",    "name": "Clementina DuBuque",    "id": 10  }]

      userData.forEach((user, userIndex) => {
        user.photo = ''//photoData.results[userIndex].picture.thumbnail;
        chatPanel.addUser(user);
      })
      chatPanel.hideAjax();

      /*
      this.getData(url.users).then((userData) => {
        this.getData(url.photos).then((photoData) => {
          userData.forEach((user, userIndex) => {
            user.photo = photoData.results[userIndex].picture.thumbnail;
            chatPanel.addUser(user);
          })
          chatPanel.hideAjax();
        });
      });
      */
    }
  }
</script>
